package common

import "github.com/apache/arrow/go/v12/arrow"

type DFSchema struct {
	*arrow.Schema
}

func (s DFSchema) Select(projection []string) DFSchema {
	fields := make([]arrow.Field, 0)
	for _, columnName := range projection {
		field, ok := s.FieldsByName(columnName)
		if ok {
			fields = append(fields, field...)
		}
	}
	newSchema := arrow.NewSchema(fields, nil)
	return DFSchema{newSchema}
}
