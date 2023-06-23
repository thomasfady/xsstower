package types

type NotifierInformation struct {
	Name   string
	Key    string
	Config []NotifyConfig
}

type NotifyConfig struct {
	Key       string
	Sensitive bool
	Value     string
	Text      string
}

type Notifier struct {
	Config []NotifyConfig
}
