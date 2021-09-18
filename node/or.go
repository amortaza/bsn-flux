package node

import (
	"fmt"
)

type Or struct {
	Left  Node
	Right Node

	nodeCompiler Compiler
}

func NewOr(nodeCompiler Compiler) *Or {
	return &Or{nodeCompiler: nodeCompiler}
}

func (orNode *Or) Compile() (string, error) {
	return orNode.nodeCompiler.OrCompile(orNode)
}

func (orNode *Or) Put(kid Node) error {

	if orNode.Left == nil {
		orNode.Left = kid
		return nil
	}

	if orNode.Right == nil {
		orNode.Right = kid
		return nil
	}

	return fmt.Errorf("no capacity to Put() a node inside an OR node")
}
