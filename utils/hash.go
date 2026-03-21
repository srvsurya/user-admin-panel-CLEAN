package utils

import (
	"Week_12/logger"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, email string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error("Hash generation failure", zap.String("Email", email), zap.Error(err))
	}
	return string(hash), err
}
