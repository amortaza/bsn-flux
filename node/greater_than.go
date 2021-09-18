package node

import (
	"fmt"
)

type GreaterThan struct {
	Left    Node
	Right   Node
	OrEqual bool

	nodeCompiler Compiler
}

func NewGreaterThan(nodeCompiler Compiler) *GreaterThan {
	return &GreaterThan{
		OrEqual:      false,
		nodeCompiler: nodeCompiler,
	}
}

func NewGreaterOrEqual(nodeCompiler Compiler) *GreaterThan {
	return &GreaterThan{
		OrEqual:      true,
		nodeCompiler: nodeCompiler,
	}
}

func (greaterThan *GreaterThan) Compile() (string, error) {
	return greaterThan.nodeCompiler.GreaterThanCompile(greaterThan)
}

func (greaterThan *GreaterThan) Put(kid Node) error {

	if greaterThan.Left == nil {
		greaterThan.Left = kid
		return nil
	}

	if greaterThan.Right == nil {
		greaterThan.Right = kid
		return nil
	}

	return fmt.Errorf("no capacity to Put() a node inside a GREATHER THAN node")
}
