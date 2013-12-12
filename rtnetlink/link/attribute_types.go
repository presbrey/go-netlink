package link

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "bitbucket.org/vtolstov/go-netlink"

const (
	IFLA_UNSPEC netlink.AttributeType = iota
	IFLA_ADDRESS
	IFLA_BROADCAST
	IFLA_IFNAME
	IFLA_MTU
	IFLA_LINK
	IFLA_QDISC
	IFLA_STATS
	IFLA_COST
	IFLA_PRIORITY
	IFLA_MASTER
	IFLA_WIRELESS
	IFLA_PROTINFO
	IFLA_TXQLEN
	IFLA_MAP
	IFLA_WEIGHT
	IFLA_OPERSTATE
	IFLA_LINKMODE
	IFLA_LINKINFO
	IFLA_NET_NS_PID
	IFLA_IFALIAS
	IFLA_NUM_VF
	IFLA_VFINFO_LIST
	IFLA_STATS64
	IFLA_VF_PORTS
	IFLA_PORT_SELF
	IFLA_AF_SPEC
	IFLA_GROUP
	IFLA_NET_NS_FD
	IFLA_EXT_MASK
	IFLA_PROMISCUITY
	IFLA_NUM_TX_QUEUES
	IFLA_NUM_RX_QUEUES
	IFLA_CARRIER
	IFLA_PHYS_PORT_ID
	IFLA_MAX = IFLA_PHYS_PORT_ID
)

var AttributeTypeStrings = map[netlink.AttributeType]string{
	IFLA_UNSPEC:        "IFLA_UNSPEC",
	IFLA_ADDRESS:       "IFLA_ADDRESS",
	IFLA_BROADCAST:     "IFLA_BROADCAST",
	IFLA_IFNAME:        "IFLA_IFNAME",
	IFLA_MTU:           "IFLA_MTU",
	IFLA_LINK:          "IFLA_LINK",
	IFLA_QDISC:         "IFLA_QDISC",
	IFLA_STATS:         "IFLA_STATS",
	IFLA_COST:          "IFLA_COST",
	IFLA_PRIORITY:      "IFLA_PRIORITY",
	IFLA_MASTER:        "IFLA_MASTER",
	IFLA_WIRELESS:      "IFLA_WIRELESS",
	IFLA_PROTINFO:      "IFLA_PROTINFO",
	IFLA_TXQLEN:        "IFLA_TXQLEN",
	IFLA_MAP:           "IFLA_MAP",
	IFLA_WEIGHT:        "IFLA_WEIGHT",
	IFLA_OPERSTATE:     "IFLA_OPERSTATE",
	IFLA_LINKMODE:      "IFLA_LINKMODE",
	IFLA_LINKINFO:      "IFLA_LINKINFO",
	IFLA_NET_NS_PID:    "IFLA_NET_NS_PID",
	IFLA_IFALIAS:       "IFLA_IFALIAS",
	IFLA_NUM_VF:        "IFLA_NUM_VF",
	IFLA_VFINFO_LIST:   "IFLA_VFINFO_LIST",
	IFLA_STATS64:       "IFLA_STATS64",
	IFLA_VF_PORTS:      "IFLA_VF_PORTS",
	IFLA_PORT_SELF:     "IFLA_PORT_SELF",
	IFLA_AF_SPEC:       "IFLA_AF_SPEC",
	IFLA_GROUP:         "IFLA_GROUP",
	IFLA_NET_NS_FD:     "IFLA_NET_NS_FD",
	IFLA_EXT_MASK:      "IFLA_EXT_MASK",
	IFLA_PROMISCUITY:   "IFLA_PROMISCUITY",
	IFLA_NUM_TX_QUEUES: "IFLA_NUM_TX_QUEUES",
	IFLA_NUM_RX_QUEUES: "IFLA_NUM_RX_QUEUES",
	IFLA_CARRIER:       "IFLA_CARRIER",
	IFLA_PHYS_PORT_ID:  "IFLA_PHYS_PORT_ID",
}

const (
	RTEXT_FILTER_VF     = 1 << iota
	RTEXT_FILTER_BRVLAN = 1 << iota
)
