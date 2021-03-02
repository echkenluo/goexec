package main

import (
  "fmt"
  "github.com/echkenluo/goexec/pkg/eventimpl"
  "github.com/echkenluo/goexec/pkg/eventmonitor"
)


func main() {
  portIpMap := make(map[uint32]string)

  portIpMap[1] = "11111"
  portIpMap[2] = "22222"

  el := eventimpl.NewEventImpl("el")
  em := eventmonitor.NewEventMonitor(el)

  em.ipMonitor.RegisterCallBack(el.updateIp(portIpMap))
  fmt.Println("main")

}
