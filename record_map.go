package flux

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/amortaza/bsn/flux/relation"
	"github.com/amortaza/bsn/logger"
)

type RecordMap struct {
	Data map[string]interface{}
}

func (recmap *RecordMap) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	datamap := recmap.Data

	// we sneak in "id" because React-Admin minimally requires "id"
	datamap["id"] = datamap["x_pk"]

	first := true
	for key, valuebytes := range datamap {

		if first {
			first = false
		} else {
			buffer.WriteString(",")
		}

		value := string(valuebytes.([]byte))

		valuemarshalled, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}

		asStr := fmt.Sprintf("\"%s\" : %s", key, string(valuemarshalled))

		buffer.WriteString(asStr)
	}

	buffer.WriteString("}")

	logger.Log(string(buffer.Bytes()), logger.JSONencoded)

	return buffer.Bytes(), nil
}

func NewRecordMap() *RecordMap {
	return &RecordMap{
		Data: make(map[string]interface{}),
	}
}

func (m *RecordMap) Put(key string, value interface{}) {
	m.Data[key] = value
}

func (m *RecordMap) Get(key string) (string, error) {
	if !m.Has(key) {
		return "", fmt.Errorf("key not '%s' not found in map", key)
	}

	asByteArray, ok := m.Data[key].([]byte)
	if !ok {
		return "", fmt.Errorf("value for key'%s' not a string", key)
	}

	return string(asByteArray), nil
}

func (m *RecordMap) GetNumber(key string) (float32, error) {
	if !m.Has(key) {
		return 0, fmt.Errorf("key not '%s' not found in map", key)
	}

	value, ok := m.Data[key].(float32)
	if !ok {
		return 0, fmt.Errorf("value for key'%s' not a number", key)
	}

	return value, nil
}

func (m *RecordMap) GetBool(key string) (bool, error) {
	if !m.Has(key) {
		return false, fmt.Errorf("key not '%s' not found in map", key)
	}

	value, ok := m.Data[key].(bool)
	if !ok {
		return false, fmt.Errorf("value for key'%s' not a bool", key)
	}

	return value, nil
}

func (m *RecordMap) Type(key string) (relation.FieldType, error) {
	if !m.Has(key) {
		return "", fmt.Errorf("key not '%s' not found in map", key)
	}

	value := m.Data[key]

	if _, ok := value.(string); ok {
		return relation.String, nil
	}

	if _, ok := value.(float32); ok {
		return relation.Number, nil
	}

	if _, ok := value.(bool); ok {
		return relation.Bool, nil
	}

	return "", fmt.Errorf("cannot determine type of value for key '%s', see '%v'", key, value)
}

func (m *RecordMap) Has(key string) bool {
	_, ok := m.Data[key]

	return ok
}

func (m *RecordMap) Combine(other *RecordMap) *RecordMap {
	result := NewRecordMap()

	for k, v := range m.Data {
		result.Put(k, v)
	}

	for k, v := range other.Data {
		result.Put(k, v)
	}

	return result
}
