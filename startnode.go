package main

import (
	"fmt"
	"log"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("启动节点 %s\n", nodeID)
	if len(minerAddress) > 0 {
		if ValidateAddress(minerAddress) {
			fmt.Println("启动挖矿节点. 接受奖励的地址为: ", minerAddress)
		} else {
			log.Panic("地址错误!")
		}
	}
	StartServer(nodeID, minerAddress)
}
