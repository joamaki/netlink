package main

import "github.com/vishvananda/netlink"

/*
sudo ip netns add foo
sudo ip link add numrxqueues 2 nk type netkit single
cli.py --spec /home/darkstar/net/Documentation/netlink/specs/netdev.yaml
--do bind-queue  --json "{\"src-ifindex\": $(cat /sys/class/net/$netdev/ifindex), \"src-queue-id\": $queue,
\"dst-ifindex\": $(cat /sys/class/net/nk/ifindex), \"queue-type\": \"rx\"}"
*/

func main() {
	// Add a single-pairing netkit device with 2 rx queues
	err := netlink.LinkAdd(&netlink.Netkit{
		LinkAttrs: netlink.LinkAttrs{
			Name:        "nk",
			NumRxQueues: 2,
		},
		Mode:    netlink.NETKIT_MODE_L2,
		Pairing: netlink.NETKIT_DEVICE_SINGLE,
	})
	if err != nil {
		panic(err)
	}

	nkLink, err := netlink.LinkByName("nk")
	if err != nil {
		panic(err)
	}

	enLink, err := netlink.LinkByName("enp2s0f1np1")
	if err != nil {
		panic(err)
	}

	// Bind the 15th rx queue of enp2s0f1np1 to the 0th queue of nk device
	err = netlink.NetdevBindQueue(
		netlink.NETDEV_QUEUE_TYPE_RX,
		enLink,
		15,
		nkLink,
	)
	if err != nil {
		panic(err)
	}
}
