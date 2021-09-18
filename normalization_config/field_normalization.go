package normalization_config

const prefix = "x_"

const PrimaryKey_FieldName = prefix + "pk"

func NormalizeRelationName(name string) string {
	return prefix + name
}

func NormalizeFieldName(fieldname string) string {
	return prefix + fieldname
}
