package flux

import "github.com/amortaza/bsn-flux/relation"

type SchemaJournalist interface {
	CreateRelation(name string) error
	DeleteRelation(name string) error
	CreateField(relationName string, field *relation.Field) error
	DeleteField(relationName string, fieldname string) error
}

type EmptySchemaJournalist struct {
}

func (empty *EmptySchemaJournalist) CreateRelation(name string) error {
	return nil
}

func (empty *EmptySchemaJournalist) DeleteRelation(name string) error {
	return nil
}

func (empty *EmptySchemaJournalist) CreateField(relationName string, field *relation.Field) error {
	return nil
}

func (empty *EmptySchemaJournalist) DeleteField(relationName string, fieldname string) error {
	return nil
}
