package utils

import (
	"errors"
	"hackathon-api/models"
	_ "os"

	"github.com/rs/zerolog/log"
)

func Checkb(b bool, msg string) {
	if !b {
		log.Error().Msgf("h => %v ", msg)
	}
}

func ValidateMoneyType(moneyType string) error {

	if _, ok := models.GetMoney()[moneyType]; ok {
		return nil
	}

	return errors.New("Invalid money type")
}
