package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haquenafeem/anonymous/internal"
	"github.com/haquenafeem/anonymous/model"
)

func (svc *Service) PostMessage(req *model.PostMessageRequest) *model.PostMessageResponse {
	user, err := svc.repo.FindUser(req.UserID)
	if err != nil || user.ID == "" {
		return &model.PostMessageResponse{
			Err: "user not found",
		}
	}

	message := &model.Message{
		ID:     internal.GenerateUUID(),
		UserID: user.ID,
		Data:   req.Message,
	}

	err = svc.repo.CreateMessage(message)
	if err != nil {
		return &model.PostMessageResponse{
			Err: "message posting failed",
		}
	}

	return &model.PostMessageResponse{
		IsSuccess: true,
	}
}

func (svc *Service) GetAllMessages(ctx *gin.Context) (int, *model.GetAllMessagesResponse) {
	claims_any, ok := ctx.Get("claims")
	if !ok {
		return http.StatusInternalServerError, &model.GetAllMessagesResponse{
			Err: "unable to get user claims",
		}
	}

	claims, ok := claims_any.(*internal.CustomClaims)
	if !ok {
		return http.StatusInternalServerError, &model.GetAllMessagesResponse{
			Err: "unable to cast user claims",
		}
	}

	messages, err := svc.repo.GetAll(claims.UserID)
	if err != nil {
		return http.StatusInternalServerError, &model.GetAllMessagesResponse{
			Err: "unable to get user messages",
		}
	}

	return http.StatusOK, &model.GetAllMessagesResponse{
		IsSuccess: true,
		Messages:  messages,
	}
}
