package tabulator

import "fmt"

type IColumn interface {
	GetName() string
	GetFieldLength() int
	GetPadLength() int
	GetFormattedValue(value any) string
	GetPaddedValue(value string) string
	SetFieldLength(length int)
}

type column struct {
	name        string
	fieldLength int
	padLength   int
}

type StringColumn struct{ column }
type IntColumn struct{ column }
type FloatColumn struct{ column }
type CurrencyColumn struct{ column }
type PercentageColumn struct{ column }

func (col *column) GetName() string {
	return col.name
}

func (col *column) GetFieldLength() int {
	return col.fieldLength
}

func (col *column) SetFieldLength(length int) {
	col.fieldLength = length
}

func (col *column) GetPadLength() int {
	return col.padLength
}

func (col *column) GetPaddedValue(value string) string {
	formatStr := fmt.Sprintf("%%-%ds", col.fieldLength+col.padLength)
	return fmt.Sprintf(formatStr, value)
}

func (col *StringColumn) GetFormattedValue(value any) string {
	return fmt.Sprintf("%s", value)
}

func (col *IntColumn) GetFormattedValue(value any) string {
	return fmt.Sprintf("%d", value)
}

func (col *FloatColumn) GetFormattedValue(value any) string {
	return fmt.Sprintf("%f", value)
}

func (col *CurrencyColumn) GetFormattedValue(value any) string {
	return fmt.Sprintf("$%.2f", value)
}

func (col *PercentageColumn) GetFormattedValue(value any) string {
	return fmt.Sprintf("%.2f%%", value.(float64)*100)
}

func newColumn(name string) column {
	return column{
		name:        name,
		fieldLength: len(name),
		padLength:   4,
	}
}

func NewStringColumn(name string) IColumn {
	return &StringColumn{newColumn(name)}
}

func NewIntColumn(name string) IColumn {
	return &IntColumn{newColumn(name)}
}

func NewFloatColumn(name string) IColumn {
	return &FloatColumn{newColumn(name)}
}

func NewCurrencyColumn(name string) IColumn {
	return &CurrencyColumn{newColumn(name)}
}

func NewPercentageColumn(name string) IColumn {
	return &PercentageColumn{newColumn(name)}
}
