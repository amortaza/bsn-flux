package node

import (
	"fmt"
)

type IsNull struct {
	ColumnNode *Column

	Not bool

	nodeCompiler Compiler
}

func NewIsNull(columnNode *Column, nodeCompiler Compiler) Node {
	return &IsNull{
		ColumnNode:   columnNode,
		nodeCompiler: nodeCompiler,
	}
}

func NewNotIsNull(columnNode *Column, nodeCompiler Compiler) Node {
	return &IsNull{
		ColumnNode:   columnNode,
		Not:          true,
		nodeCompiler: nodeCompiler,
	}
}

func (isNull *IsNull) Compile() (string, error) {
	return isNull.nodeCompiler.IsNullCompile(isNull)
}

func (isNull *IsNull) Put(kid Node) error {
	return fmt.Errorf("no capacity to Put() a node inside an IS NULL node")
}
