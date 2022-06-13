package tabulator

import (
	"strings"
)

// Tabulator definition

type Tabulator struct {
	columns []IColumn
	nameMap map[string]int
	values  [][]string
	count   int
}

func NewTabulator(columnNames ...string) *Tabulator {
	t := &Tabulator{
		columns: make([]IColumn, 0),
		nameMap: make(map[string]int),
		values:  make([][]string, 0),
		count:   0,
	}

	for _, name := range columnNames {
		t.nameMap[name] = len(t.columns)
		t.columns = append(t.columns, NewStringColumn(name))
	}

	return t
}

func (tab *Tabulator) AddColumn(column IColumn) {
	tab.nameMap[column.GetName()] = len(tab.columns)
	tab.columns = append(tab.columns, column)
}

func (tab *Tabulator) AddRow() {
	tab.values = append(tab.values, make([]string, len(tab.columns)))
	tab.count += 1
}

func (tab *Tabulator) updateFieldLength(column int, value any) {
	v := tab.columns[column].GetFormattedValue(value)
	if len(v) > tab.columns[column].GetFieldLength() {
		tab.columns[column].SetFieldLength(len(v))
	}
}

func (tab *Tabulator) AddValueByColumnIndex(row int, column int, value any) {
	tab.updateFieldLength(column, value)
	tab.values[row][column] = tab.columns[column].GetFormattedValue(value)
}

func (tab *Tabulator) AddValueByColumnName(row int, column string, value any) {
	idx := tab.nameMap[column]
	tab.updateFieldLength(idx, value)
	tab.values[row][idx] = tab.columns[idx].GetFormattedValue(value)
}

func (tab *Tabulator) GetTableHeader() string {
	header := strings.Builder{}
	header.WriteString("\n")
	for _, col := range tab.columns {
		header.WriteString(col.GetPaddedValue(col.GetName()))
	}
	header.WriteString("\n")
	for _, col := range tab.columns {
		header.WriteString(strings.Repeat("-", col.GetFieldLength()+col.GetPadLength()))
	}

	return header.String()
}

func (tab *Tabulator) ToRow(record []string) string {
	row := strings.Builder{}
	for i, col := range tab.columns {
		row.WriteString(col.GetPaddedValue(record[i]))
	}
	return row.String()
}

func (tab *Tabulator) ToTable() string {
	output := strings.Builder{}
	output.WriteString(tab.GetTableHeader() + "\n")

	if tab.count == 0 {
		output.WriteString("No records found\n")
	}

	// write the records to the string builder
	for _, record := range tab.values {
		output.WriteString(tab.ToRow(record) + "\n")
	}

	return output.String()
}

func (tab *Tabulator) ToSegmentedTable(segmentIndex int) string {
	output := strings.Builder{}
	output.WriteString(tab.GetTableHeader() + "\n")

	if tab.count == 0 {
		output.WriteString("No records found\n")
	}

	segmentValue := ""
	isValidSegmentation := false
	if segmentIndex >= 0 && segmentIndex < len(tab.values[0]) {
		segmentValue = tab.values[0][segmentIndex]
		isValidSegmentation = true
	}

	for _, record := range tab.values {
		if isValidSegmentation && record[segmentIndex] != segmentValue {
			output.WriteString(tab.GetTableHeader() + "\n")
			segmentValue = record[segmentIndex]
		}

		output.WriteString(tab.ToRow(record) + "\n")
	}

	return output.String()
}

func (tab *Tabulator) calculateFieldLengths() {
	for _, fields := range tab.values {
		for i := 0; i < len(fields); i++ {
			if len(fields[i]) > tab.columns[i].GetFieldLength() {
				tab.columns[i].SetFieldLength(len(fields[i]))
			}
		}
	}
}
