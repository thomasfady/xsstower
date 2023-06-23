package notify

import (
	"time"

	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/notify/types"
)

type NotifyConfig struct {
	Key       string
	Sensitive bool
	Value     string
}

type NotifierInterface interface {
	Notify(hit models.XssHit) error
	GetBaseInformations() (infos types.NotifierInformation)
}

type Notifier struct {
	Config []NotifyConfig
}

func (n *Notifier) GetConfig(key string) (value string) {
	for _, v := range n.Config {
		if v.Key == key {
			value = v.Value
			return
		}
	}
	value = ""
	return
}

func NotifiersList() (ni []NotifierInterface) {
	t := &TelegramNotifier{}
	ni = append(ni, t)
	return
}



func NotifyHit(correlation_key string) {
	time.Sleep(10 * time.Second)

	var hit models.XssHit
	models.DB.Preload("Handler").First(&hit, "correlation_key = ?", correlation_key)

	t := &TelegramNotifier{}
	t.Config = []NotifyConfig{
		{Value: "CHANGEME", Sensitive: false, Key: "chat_id"},
		{Value: "CHANGEME", Sensitive: true, Key: "token"},
	}
	
	t.Notify(hit)
}
