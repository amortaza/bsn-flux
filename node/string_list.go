package node

import (
	"fmt"
)

type StringList struct {
	Elements []string

	nodeCompiler Compiler
}

func NewStringList(elements []string, nodeCompiler Compiler) Node {
	return &StringList{
		Elements:     elements,
		nodeCompiler: nodeCompiler,
	}
}

func (stringList *StringList) Compile() (string, error) {
	return stringList.nodeCompiler.StringListCompile(stringList)
}

func (stringList *StringList) Put(kid Node) error {
	return fmt.Errorf("no capacity to Put() a node inside a LIST OF STRINGS node")
}
