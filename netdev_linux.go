package netlink

import (
	"github.com/vishvananda/netlink/nl"
	"golang.org/x/sys/unix"
)

type NetdevQueueType uint32

const (
	NETDEV_QUEUE_TYPE_RX NetdevQueueType = iota
	NETDEV_QUEUE_TYPE_TX
)

func (h *Handle) NetdevBindQueue(queueType NetdevQueueType, srcLink Link, srcQueueId uint32, dstLink Link) error {
	f, err := h.GenlFamilyGet(nl.GENL_NETDEV_NAME)
	if err != nil {
		return err
	}
	msg := &nl.Genlmsg{
		Command: nl.GENL_NETDEV_CMD_BIND_QUEUE,
		Version: nl.GENL_NETDEV_VERSION,
	}
	req := h.newNetlinkRequest(int(f.ID), 0)
	req.AddData(msg)
	req.AddData(nl.NewRtAttr(nl.GENL_NETDEV_ATTR_QUEUE_PAIR_QUEUE_TYPE, nl.Uint32Attr(uint32(queueType))))
	req.AddData(nl.NewRtAttr(nl.GENL_NETDEV_ATTR_QUEUE_PAIR_SRC_IFINDEX, nl.Uint32Attr(uint32(srcLink.Attrs().Index))))
	req.AddData(nl.NewRtAttr(nl.GENL_NETDEV_ATTR_QUEUE_PAIR_SRC_QUEUE_ID, nl.Uint32Attr(srcQueueId)))
	req.AddData(nl.NewRtAttr(nl.GENL_NETDEV_ATTR_QUEUE_PAIR_DST_IFINDEX, nl.Uint32Attr(uint32(dstLink.Attrs().Index))))

	_, err = req.Execute(unix.NETLINK_GENERIC, 0)
	return err
}

func NetdevBindQueue(queueType NetdevQueueType, srcLink Link, srcQueueId uint32, dstLink Link) error {
	return pkgHandle.NetdevBindQueue(queueType, srcLink, srcQueueId, dstLink)
}
