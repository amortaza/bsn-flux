package node

import (
	"fmt"
)

type Equal struct {
	Left  Node
	Right Node
	Not   bool

	nodeCompiler Compiler
}

func NewEqual(nodeCompiler Compiler) *Equal {
	return &Equal{nodeCompiler: nodeCompiler}
}

func NewNotEqual(nodeCompiler Compiler) *Equal {
	return &Equal{
		Not:          true,
		nodeCompiler: nodeCompiler,
	}
}

func (equal *Equal) Compile() (string, error) {
	return equal.nodeCompiler.EqualCompile(equal)
}

func (equal *Equal) Put(kid Node) error {
	if equal.Left == nil {
		equal.Left = kid
		return nil
	}

	if equal.Right == nil {
		equal.Right = kid
		return nil
	}

	return fmt.Errorf("no capacity to Put() a node inside an EQUAL node")
}
