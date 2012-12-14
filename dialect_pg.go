package hood

import (
	"fmt"
	"strings"
)

type DialectPg struct{}

func (d *DialectPg) Name() string {
	return "postgres"
}

func (d *DialectPg) Pk() string {
	return "id"
}

func (d *DialectPg) Quote(s string) string {
	q := `"`
	return strings.Join([]string{q, s, q}, "")
}

func (d *DialectPg) MarkerStartPos() int {
	return 1
}

func (d *DialectPg) Marker(pos int) string {
	return fmt.Sprintf("$%d", pos)
}

func (d *DialectPg) SqlType(f interface{}, size int, autoIncr bool) string {
	switch f.(type) {
	case bool:
		return "boolean"
	case int, int8, int16, int32, uint, uint8, uint16, uint32:
		if autoIncr {
			return "serial"
		}
		return "integer"
	case int64, uint64:
		if autoIncr {
			return "bigserial"
		}
		return "bigint"
	case float32, float64:
		return "double precision"
	case []byte:
		return "bytea"
	case string:
		return "text"
	}
	if size < 1 {
		size = 255
	}
	return fmt.Sprintf("varchar(%d)", size)
}
