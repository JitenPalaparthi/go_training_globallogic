package messaging

import "context"

type IMessage interface {
	Publish(context.Context) error
	Subscribe(context.Context) error
}
