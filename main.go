package main

import (
	"log"
	"net"

	"github.com/yianding/net-android/anet"
)

func RawInterface() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
	}

	for _, i := range interfaces {
		log.Println(i)
	}
}
func AnetInterface() {
	interfaces, err := anet.Interfaces()
	if err != nil {
		log.Println(err)
	}

	for _, i := range interfaces {
		log.Println(i)
	}
}
func NetInterfaceAddrs() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
	}

	for _, addr := range addrs {
		log.Println(addr)
	}
}
func AnetInterfaceAddrs() {
	addrs, err := anet.InterfaceAddrs()
	if err != nil {
		log.Println(err)
	}

	for _, addr := range addrs {
		log.Println(addr)
	}
}
func main() {
	anet.SetAndroidVersion(11)
	RawInterface()
	log.Println("----------------")
	AnetInterface()
	log.Println("----------------")
	NetInterfaceAddrs()
	log.Println("----------------")
	AnetInterfaceAddrs()
}
