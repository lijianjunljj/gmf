package wrappers

import (
	"github.com/lijianjunljj/gmfcommon/wrapper"
	"github.com/micro/go-micro/v2/client"
)

type userWrapper struct {
	*wrapper.Wrapper
}

func NewUserWrapper(c client.Client) client.Client {
	w := new(userWrapper)
	w.Wrapper = wrapper.NewWrapper(c)
	return w
}
