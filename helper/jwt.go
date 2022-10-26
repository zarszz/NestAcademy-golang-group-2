package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zarszz/NestAcademy-golang-group-2/config"
)

type Token struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}

func CreateToken(payload *Token) (string, error) {
	claims := jwt.MapClaims{
		"payload": payload,
		"issued":  time.Now().Add(10 * time.Minute),
	}
	var conf, _ = config.LoadConfig(".")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokString string) (*Token, error) {
	var conf, _ = config.LoadConfig(".")
	tok, err := jwt.Parse(tokString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(conf.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !tok.Valid {
		return nil, errors.New("invalid token")

	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")

	}

	// TODO : verify issued token
	issuedString := fmt.Sprintf("%v", claims["issued"])
	issued, err := time.Parse(time.RFC3339, issuedString)
	if err != nil {
		return nil, err
	}

	if time.Now().After(issued) {
		return nil, errors.New("token expired")
	}

	byteClaims, err := json.Marshal(claims["payload"])
	if err != nil {
		return nil, err
	}

	var myToken Token
	err = json.Unmarshal(byteClaims, &myToken)
	if err != nil {
		return nil, err
	}

	return &myToken, nil

}
