package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github/jailtonjunior94/go-telegram/src/dtos"
	"github/jailtonjunior94/go-telegram/src/environments"
	"net/http"
)

type TelegramService struct{}

func NewTelegramService() TelegramService {
	return TelegramService{}
}

func (t TelegramService) SendMessage(message string) error {
	request := dtos.SendMessageReqBody{
		ChatID: int64(environments.ChatID),
		Text:   message,
	}

	reqBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s/sendMessage", environments.TelegramBaseUrl, environments.BotKey)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
