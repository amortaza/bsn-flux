package node

import (
	"fmt"
)

type Number struct {
	Text string

	nodeCompiler Compiler
}

func NewNumber(value float32, nodeCompiler Compiler) Node {

	return &Number{
		Text:         fmt.Sprintf("%f", value),
		nodeCompiler: nodeCompiler,
	}
}

func (number *Number) Compile() (string, error) {
	return number.nodeCompiler.NumberCompile(number)
}

func (number *Number) Put(kid Node) error {
	return fmt.Errorf("no capacity to Put() a node inside a NUMBER node")
}
