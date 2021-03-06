package main

/* The output of this utility is JSON, but not intended to be easily human-readable.
   At this moment it exists to test the rtnetlink/link subsystem */

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "github.com/vishvananda/go-netlink/rtnetlink/link"
import "github.com/vishvananda/go-netlink/rtnetlink"
import "log"
import "github.com/vishvananda/go-netlink"
import "encoding/binary"

func main() {
	// This imitates the functionality of rtnl_wilddump_req_filter from bridge/vlan.c in the
	// iproute2 code. Note that the filter requires kernel 3.9 or newer
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, link.RTEXT_FILTER_BRVLAN)
	lnmsg := link.NewHeader(rtnetlink.AF_BRIDGE, 0, 0, 0, 0)
	msg := rtnetlink.NewMessage(lnmsg, []netlink.Attribute{ {link.IFLA_EXT_MASK, data} })
	nlmsg, err := netlink.NewMessage(rtnetlink.RTM_GETLINK, netlink.NLM_F_DUMP|netlink.NLM_F_REQUEST, msg)
	if err != nil {
		log.Panicf("Couldn't construct message: %v", err)
	}
	nlsock, err := netlink.Dial(netlink.NETLINK_ROUTE)
	if err != nil {
		log.Panicf("Couldn't dial netlink: %v", err)
	}
	h := netlink.NewHandler(nlsock)
	ec := make(chan error)
	go h.Start(ec)
	c, err := h.Query(*nlmsg, 1)
	if err != nil {
		log.Panicf("Couldn't write netlink: %v", err)
	}
	for i := range c {
		if i.Header.MessageType() == netlink.NLMSG_DONE {
			break
		}
		switch i.Header.MessageType() {
		case rtnetlink.RTM_NEWLINK:
			hdr := &link.Header{}
			msg := rtnetlink.NewMessage(hdr, nil)
			err = msg.UnmarshalNetlink(i.Body)
			if err == nil {
				log.Printf("Link[%d] (Family: %v; Type: %v; Flags: %v; Changes: %v)",
					hdr.InterfaceIndex(),
					hdr.InterfaceFamily(), hdr.InterfaceType(), hdr.Flags(),
					hdr.InterfaceChanges())
				for i := range msg.Attributes {
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
