package actions

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"time"

	"github.com/SrLeet03/decentify-backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

func Login(auth models.LoginReq) models.LoginResponse {
	var resp models.LoginResponse

	err, token := generateJWT(resp.Userid)
	if err != "" {
		resp.Error = err
		return resp
	}
	resp.Token = token
	resp.Error = ""
	return resp
}

func generateJWT(id string) (string, string) {

	token, err := CreateToken(id)

	if err != "" {
		return "Server error", ""
	}

	return "", token
}

func CreateToken(userId string) (string, string) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true

	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 256).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("SECRET_KEY_JWT_TOKEN")))
	if err != nil {
		return "", "Server error"
	}
	return token, ""
}

func ValidateToken(tokenString string) (string, string) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.JwtSchema{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY_JWT_TOKEN")), nil
		},
	)
	if err != nil {
		return "", "invali_token"
	}
	claims, ok := token.Claims.(*models.JwtSchema)
	if !ok {
		return "", "invali_token"
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return "", "token_expired"
	}
	return "valid_token", ""
}

func GenerateUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := uuidWithHyphen.String()

	return uuid
}

func ValidatePassword(password string, original string) string {

	if password != original {
		return "Inavlid password"
	}

	return ""
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
