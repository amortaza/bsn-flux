package normalization

const prefix = "x_"
const PrimaryKeyFieldname = prefix + "pk"

func NormalizeRelationName(name string) string {
	return prefix + name
}

func NormalizeFieldName(fieldname string) string {
	return prefix + fieldname
}
