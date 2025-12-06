package utils

import (
	"DB_LAB/config"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(userID uuid.UUID, isAdmin bool) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":  userID.String(),
			"is_admin": isAdmin,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(config.Env.Secret))
	if err != nil {
		return "", errors.New(fmt.Sprintf("Ошибка генерации токена, повторите попытку: %v", err))
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string) (userID uuid.UUID, isAdmin bool, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Env.Secret), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, ok1 := claims["user_id"].(string)
		isAdmin, ok2 := claims["is_admin"].(bool)
		if ok1 && ok2 {
			userUUID, err := uuid.Parse(userID)
			if err != nil {
				return uuid.Nil, false, err
			}

			return userUUID, isAdmin, nil
		}
	}
	return uuid.Nil, false, errors.New("invalid token")
}
