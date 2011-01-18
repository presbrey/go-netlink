package main

/* The output of this utility is JSON, but not intended to be easily human-readable.
   At this moment it exists to test the rtnetlink/route subsystem */

import "os"
import "netlink/rtnetlink/link"
import "netlink/rtnetlink"
import "log"
import "netlink"

func main(){
  nlmsg, err := netlink.NewMessage(rtnetlink.RTM_GETLINK, netlink.NLM_F_DUMP|netlink.NLM_F_REQUEST, &link.Header{},4)
  if err != nil {
    log.Exitf("Couldn't construct message: %v", err)
  }
  nlsock, err := netlink.Dial(netlink.NETLINK_ROUTE)
  if err != nil {
    log.Exitf("Couldn't dial netlink: %v", err)
  }
  h := netlink.NewHandler(nlsock)
  ec := make(chan os.Error)
  go h.Start(ec)
  c := make(chan netlink.Message)
  err = h.SendQuery(*nlmsg, c)
  if err != nil {
    log.Exitf("Couldn't write netlink: %v", err)
  }
  for i := range( c) {
    switch i.Header.MessageType() {
      case rtnetlink.RTM_NEWLINK:
        msg := rtnetlink.NewMessage(&link.Header{}, nil)
        err = msg.UnmarshalNetlink(i.Body, 4)
        if err == nil {
          for i := range(msg.Attributes){
            log.Printf("Attribute[%d]: %v", i, msg.Attributes[i])
          }
        } else {
          log.Printf("Unmarshal error: %v", err)
        }
      default:
          log.Printf("Unknown type: %v", i)
    }
  }
}