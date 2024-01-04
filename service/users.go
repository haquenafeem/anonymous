package service

import (
	"errors"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/haquenafeem/anonymous/internal"
	"github.com/haquenafeem/anonymous/model"
)

func (svc *Service) RegisterUser(req *model.RegisterRequest) *model.RegisterResponse {
	if req == nil {
		return &model.RegisterResponse{
			Err: "request body is nil",
		}
	}

	existingUser, err := svc.repo.FindByEmail(req.Email)
	if err != nil {
		return &model.RegisterResponse{
			Err: "user creation failed",
		}
	}

	if existingUser.ID != "" {
		return &model.RegisterResponse{
			Err: "user already registered with email",
		}
	}

	hashedPassword, err := internal.HashPassword(req.Password)
	if err != nil {
		return &model.RegisterResponse{
			Err: "hashing failed",
		}
	}

	user := &model.User{
		ID:       internal.GenerateUUID(),
		Email:    req.Email,
		Password: hashedPassword,
	}

	err = svc.repo.CreateUser(user)
	if err != nil {
		return &model.RegisterResponse{
			Err: "user creation failed",
		}
	}

	return &model.RegisterResponse{
		IsSuccess: true,
	}
}

func (svc *Service) LoginUser(req *model.LoginRequest) *model.LoginResponse {
	if req == nil {
		return &model.LoginResponse{
			Err: "request body is nil",
		}
	}

	user, err := svc.repo.FindByEmail(req.Email)
	if err != nil {
		return &model.LoginResponse{
			Err: "error fetching user from db",
		}
	}

	err = internal.ComparePasswords(user.Password, req.Password)
	if err != nil {
		return &model.LoginResponse{
			Err: "password mismatch",
		}
	}

	token, err := internal.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return &model.LoginResponse{
			Err: "token generation failed",
		}
	}

	return &model.LoginResponse{
		IsSuccess: true,
		Token:     token,
	}
}

func (svc *Service) UploadProfilePic(ctx *gin.Context) (int, *model.UploadResponse) {
	file, err := ctx.FormFile("upload_file")

	if err != nil {
		return http.StatusBadRequest, &model.UploadResponse{
			Err: "no file recieved",
		}
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	if err := ctx.SaveUploadedFile(file, "./images/"+newFileName); err != nil {
		return http.StatusInternalServerError, &model.UploadResponse{
			Err: "unable to save the file",
		}
	}

	claims_any, ok := ctx.Get("claims")
	if !ok {
		return http.StatusInternalServerError, &model.UploadResponse{
			Err: "unable to get user claims",
		}
	}

	claims, ok := claims_any.(*internal.CustomClaims)
	if !ok {
		return http.StatusInternalServerError, &model.UploadResponse{
			Err: "unable to cast user claims",
		}
	}

	res := svc.repo.UpdateProfilePicId(claims.UserID, newFileName)
	if !res.IsSuccess {
		return http.StatusInternalServerError, res
	}

	return http.StatusOK, res
}

func (svc *Service) ProfilePic(ctx *gin.Context) (int, *model.ProfilePicResponse) {
	claims_any, ok := ctx.Get("claims")
	if !ok {
		return http.StatusInternalServerError, &model.ProfilePicResponse{
			Err: "unable to get user claims",
		}
	}

	claims, ok := claims_any.(*internal.CustomClaims)
	if !ok {
		return http.StatusInternalServerError, &model.ProfilePicResponse{
			Err: "unable to cast user claims",
		}
	}

	user, err := svc.repo.FindUser(claims.UserID)
	if err != nil || user.ID == "" {
		return http.StatusInternalServerError, &model.ProfilePicResponse{
			Err: "user not found",
		}
	}

	if user.Profile_Pic_ID == "" {
		return http.StatusInternalServerError, &model.ProfilePicResponse{
			Err: "profile pic not found",
		}
	}

	return http.StatusOK, &model.ProfilePicResponse{
		IsSuccess:    true,
		ProfilePicID: user.Profile_Pic_ID,
	}
}

func (svc *Service) GenerateQrCodeBytes(ctx *gin.Context) ([]byte, error) {
	claims_any, ok := ctx.Get("claims")
	if !ok {
		return nil, errors.New("unable to get user claims")
	}

	claims, ok := claims_any.(*internal.CustomClaims)
	if !ok {
		return nil, errors.New("unable to cast user claims")
	}

	messageUrl := ctx.Request.Host + "/messages/" + claims.UserID
	return internal.GenerateQRCode(messageUrl)
}
