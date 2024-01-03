package internal

import (
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQRCode(content string) ([]byte, error) {
	var png []byte
	png, err := qrcode.Encode(content, qrcode.Medium, 512)

	return png, err
}
