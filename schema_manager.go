package flux

import (
	"github.com/amortaza/bsn/flux/relation"
)

type SchemaManager interface {
	CreateRelation(name string) error
	DeleteRelation(name string) error
	CreateField(relationName string, field *relation.Field) error
	DeleteField(relationName string, fieldname string) error
}
