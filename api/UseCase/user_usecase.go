package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Reeeid/TodoTetris/Domain/model"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	if os.Getenv("SECRET_KEY") == "" {
		panic("SECRET_KEY is not set! Critical error")
	}
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) RegisterUser(user *model.User) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.PasswordHash = string(hashed)
	if err := u.repo.CreateUser(user); err != nil {
		return "", err
	}
	payload := map[string]interface{}{
		"username": user.Username,
	}
	token, err := GenerateJWT(payload, os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserUseCase) LoginUser(user *model.User) (string, error) {
	_, result, err := u.repo.FindByUserID(user.Username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(user.PasswordHash)); err != nil {
		return "", fmt.Errorf("Not Authenticated")
	}
	payload := map[string]interface{}{
		"username": user.Username,
	}
	token, err := GenerateJWT(payload, os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GenerateJWT(payload map[string]interface{}, secret string) (string, error) {
	//ヘッダー
	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}
	headerJSON, _ := json.Marshal(header)
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)
	payloadJSON, _ := json.Marshal(payload)
	payloadEncoded := base64.RawURLEncoding.EncodeToString(payloadJSON)
	unsignedToken := headerEncoded + "." + payloadEncoded
	//署名 (HMAC-SHA256)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedToken))
	signature := h.Sum(nil)
	signatureEncoded := base64.RawURLEncoding.EncodeToString(signature)
	return unsignedToken + "." + signatureEncoded, nil
}
