package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"

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

func (u *UserUseCase) RegisterUser(user *model.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashed)
	return u.repo.CreateUser(user)
}

func (u *UserUseCase) LoginUser(user *model.User) (string, error) {
	_, result, err := u.repo.FindByUserID(user.Username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(user.PasswordHash)); err != nil {
		return "", fmt.Errorf("Not Authenticated")
	}
	token, err := GenerateJWT(user.Username, os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GenerateJWT(username string, secret string) (string, error) {
	//ヘッダー作成
	headerRaw, _ := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	header := base64.RawURLEncoding.EncodeToString(headerRaw)
	//ペイロードの作成
	payLoadRwa, _ := json.Marshal(map[string]interface{}{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 31).Unix(),
	})
	payload := base64.RawURLEncoding.EncodeToString(payLoadRwa)
	//署名の作成
	unsignedToken := header + "." + payload
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedToken))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	//JWTの完成
	return unsignedToken + "." + signature, nil
}
