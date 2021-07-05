package bot

import (
	"encoding/json"
	"fmt"
	"github.com/Ragontar/TGBot_phoneBook/pkg/setup"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var tgBot *Bot

type Bot struct {
	token       string
	queryPrefix string
	webhook     string
}

func NewBot(tokenString string) (*Bot, error) {
	b := new(Bot)

	b.SetToken(tokenString)
	err := b.MakeQueryPrefix()
	if err != nil {
		return nil, err
	}

	tgBot = b
	return b, nil
}

func (b *Bot) SetToken(tokenString string) {
	b.token = tokenString
}

func (b *Bot) MakeQueryPrefix() error {
	if b.token == "" {
		return &TokenError{}
	}

	b.queryPrefix = fmt.Sprintf("https://api.telegram.org/bot%s/", b.token)
	return nil
}

func (b *Bot) SendMessage(chatID int64, text string) (string, error) {
	querySuffix := fmt.Sprintf("sendMessage?chat_id=%d&text=%s", chatID, text)
	query := b.queryPrefix + querySuffix

	response, err := http.PostForm(query, url.Values{})
	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		setup.GetLogger().Printf("error in parsing telegram answer %s", err)
		return "", err
	}

	return string(bodyBytes), nil
}

func (b *Bot) SetWebhook(wh string) error {
	b.webhook = wh
	querySuffix := fmt.Sprintf("setWebhook?url=%s", wh)

	resp, err := http.Get(b.queryPrefix + querySuffix)
	if err != nil {
		return err
	}

	var dec map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&dec)
	if err != nil {
		return err
	}
	if dec["ok"] == "false" {
		log.Fatal("error setting webhook")
	}

	return nil
}
