package eventimpl

import (
  "fmt"
)

type eventimpl struct {
  name string
}
func NewEventImpl(name string) *eventimpl {
  return &eventimpl{
    name: name,
  }
}
type updateIpFunc func(map[uint32]string)

func (el *eventimpl) RegisterCallBack(callback func(portIpMap map[uint32]string)) {
  return
}

func updateIp(portIpMap map[uint32]string) {
  fmt.Println("ipUpdate")
  for port, ip := range portIpMap {
  }
}
