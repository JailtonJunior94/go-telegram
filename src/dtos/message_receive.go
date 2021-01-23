package dtos

import "github/jailtonjunior94/go-telegram/src/utils"

type MessageReceive struct {
	Application string     `json:"application"`
	Method      string     `json:"method"`
	Date        utils.Time `json:"date"`
	Request     string     `json:"request"`
	Response    string     `json:"response"`
	Exception   string     `json:"exception"`
}
