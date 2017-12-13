package nonebroker

import "github.com/micro/go-micro/broker"

// 一个空的 broker
type noneBroker struct{}

func NewBroker() broker.Broker {
	return &noneBroker{}
}

func (*noneBroker) Options() broker.Options {
	return broker.Options{}
}

func (*noneBroker) Address() string {
	return ""
}

func (*noneBroker) Connect() error {
	return nil
}

func (*noneBroker) Disconnect() error {
	return nil
}

func (*noneBroker) Init(...broker.Option) error {
	return nil
}

func (*noneBroker) Publish(string, *broker.Message, ...broker.PublishOption) error {
	return nil
}

func (*noneBroker) Subscribe(string, broker.Handler, ...broker.SubscribeOption) (broker.Subscriber, error) {
	return nil, nil
}

func (*noneBroker) String() string {
	return "noneBroker"
}
