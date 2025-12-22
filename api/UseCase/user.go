package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/Reeeid/TodoTetris/Domain/model"
)

func LoginUser(user *model.LoginUser) (string, error) {
	//dbからユーザー情報取得
	username, _ := user.Username, user.Password
	//パスワード照合
	//dbだるい！
	//token発行

	token, err := GenerateJWT(username, "your-secret-key")
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
