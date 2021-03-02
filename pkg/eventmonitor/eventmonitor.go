package eventmonitor
import (
  "fmt"
)

type ipMonitor interface {
  RegisterCallBack(handler func(ofports map[uint32]string))
}
type eventMonitor struct {
  name string
  ipMap map[uint32]string
  ipMonitor ipMonitor
}

func NewEventMonitor(ipMon ipMonitor) *eventMonitor {
  ipMap := make(map[uint32]string)
  return &eventMonitor{
    ipMonitor: ipMon,
    ipMap: ipMap,
  }
}

func (e *eventMonitor) handleIpUpdate(portIpMap map[uint32]string) {
  for port, ip := range portIpMap {
    e.ipMap[port] = ip
    fmt.Println(ip)
  }
}
