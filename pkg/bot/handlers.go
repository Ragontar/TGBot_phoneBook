package bot

import (
	"github.com/Ragontar/TGBot_phoneBook/pkg/setup"
	"net/http"
	"net/url"
)

func HandleTelegramWebhook(w http.ResponseWriter, r *http.Request) {
	update, err := parseTelegramRequest(r)
	if err != nil {
		setup.GetLogger().Printf("error parsing update, %s", err.Error())
		return
	}

	text := url.PathEscape("Test Response from Server ++ $#2")
	telegramResponseBody, err := tgBot.SendMessage(update.Message.Chat.Id, text)
	if err != nil {
		setup.GetLogger().Printf("error %v during sending message, %s", err, telegramResponseBody)
		return
	}

}
