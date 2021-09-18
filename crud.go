package flux

import "github.com/amortaza/bsn/flux/node"

type CRUD interface {
	Compiler() node.Compiler

	Query(relationName string, root node.Node) error
	Next() (*RecordMap, error)

	Create(relationName string, values *RecordMap) (string, error)
	Update(relationName string, id string, values *RecordMap) error
	Delete(relationName string, id string) error
}
