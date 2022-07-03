package chat

import (
	_ "fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/multiformats/go-multiaddr"
)

func NewNode() (host.Host, error) {
	// 装载监听地址
	addr,_ :=multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/0")
	host, err := libp2p.New(
		libp2p.ListenAddrs(addr),
	)
	
	return host, err
}

// crypto 选择节点间的加密方式
func crypto() {

}
func ChooseNode() {

}
