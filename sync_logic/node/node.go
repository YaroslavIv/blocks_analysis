package node

import "fmt"

type Node interface {
	Run()
}

func Init(typeNode TypeNode, rawurl string, maxCountBlock uint64, nameTable, ramAddr, dbAddr, dbDatabase, dbUsername, dbPassword string) Node {
	var node Node

	switch typeNode {
	case ETH:
		node = InitEth(rawurl, maxCountBlock, nameTable, ramAddr, dbAddr, dbDatabase, dbUsername, dbPassword)
	default:
		panic(fmt.Sprintf("Not correct typeNode = %d", typeNode))
	}

	return node
}
