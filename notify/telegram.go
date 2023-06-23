package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/notify/types"
)

type TelegramNotifier struct {
	Notifier
}

func (n *TelegramNotifier) GetBaseInformations() (infos types.NotifierInformation) {
	infos.Name = "Telegram"
	infos.Key = "telegram"
	infos.Config = []types.NotifyConfig{
		{Sensitive: false, Key: "chat_id", Text: "Chat ID"},
		{Sensitive: true, Key: "token", Text: "Bot ID"},
	}
	return
}

func (n *TelegramNotifier) Notify(hit models.XssHit) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", n.GetConfig("token"))

	var msg string

	if hit.Url == "" {
		msg = fmt.Sprintf("\xF0\x9F\x94\xA5 Uncomplete Hit received!\nIP: %s\nHandler: %s/%s", hit.Ip, hit.Handler.Domain, hit.Handler.Path)
	} else {
		msg = fmt.Sprintf("\xF0\x9F\x94\xA5 Hit received!\nIP: %s\nUrl: %s", hit.Ip, hit.Url)
	}

	body, _ := json.Marshal(map[string]string{
		"chat_id": n.GetConfig("chat_id"),
		"text":    msg,
	})
	_, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
