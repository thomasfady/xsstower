package notify

import (
	"time"

	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/notify/types"
)

type NotifierInterface interface {
	Notify(hit models.XssHit) error
	CanNotify() bool
	GetBaseInformations() (infos types.NotifierInformation)
	SetConfig(key string, value string)
}

type Notifier struct {
	Config []types.NotifyConfig
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

func (n *Notifier) SetConfig(key string, value string) {
	for k, v := range n.Config {
		if v.Key == key {
			n.Config[k].Value = value
		}
	}
}

func NotifiersList() (ni []NotifierInterface) {
	t := &TelegramNotifier{}
	t.Config = t.GetBaseInformations().Config
	ni = append(ni, t)
	return
}



func NotifyHit(correlation_key string) {
	time.Sleep(10 * time.Second)

	var hit models.XssHit
	models.DB.Preload("Handler").Preload("Owner").First(&hit, "correlation_key = ?", correlation_key)
	for _, n := range NotifiersList() {
		infos := n.GetBaseInformations()
		for _, config := range infos.Config {
			val, ok := hit.Owner.NotifiersConfig[infos.Key+"."+config.Key]
			if ok {
				n.SetConfig(config.Key, val)
			}
		}
		if n.CanNotify() {
			n.Notify(hit)
		}
	}
}
