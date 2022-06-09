package tabulator

import (
	"fmt"
	"strings"
)

// Column definition

type Column struct {
	name        string
	fieldLength int
	padLength   int
}

func NewDefaultColumn(name string) *Column {
	return &Column{
		name:        name,
		padLength:   4,
		fieldLength: len(name),
	}
}

func NewColumn(name string, padLength int, fieldLength int) *Column {
	return &Column{
		name:        name,
		padLength:   padLength,
		fieldLength: fieldLength,
	}
}

func (c *Column) GetJustifiedValue(value string) string {
	formatStr := fmt.Sprintf("%%-%ds", c.fieldLength+c.padLength)
	return fmt.Sprintf(formatStr, value)
}

// Tabulator definition

type Tabulator struct {
	columnDefinitions []*Column
}

func NewTabulator(columnNames ...string) *Tabulator {
	t := &Tabulator{
		columnDefinitions: make([]*Column, 0),
	}

	for _, name := range columnNames {
		t.AddColumn(name)
	}

	return t
}

func (tab *Tabulator) AddColumn(name string) {
	tab.columnDefinitions = append(tab.columnDefinitions, NewDefaultColumn(name))
}

func (tab *Tabulator) AddColumnDefinition(column *Column) {
	tab.columnDefinitions = append(tab.columnDefinitions, column)
}

func (tab *Tabulator) GetTableHeader() string {
	header := strings.Builder{}
	header.WriteString("\n")
	for _, col := range tab.columnDefinitions {
		header.WriteString(col.GetJustifiedValue(col.name))
	}
	header.WriteString("\n")
	for _, col := range tab.columnDefinitions {
		header.WriteString(strings.Repeat("-", col.fieldLength+col.padLength))
	}

	return header.String()
}

func (tab *Tabulator) ToRow(record []string) string {
	row := strings.Builder{}
	for i, col := range tab.columnDefinitions {
		row.WriteString(col.GetJustifiedValue(record[i]))
	}
	return row.String()
}

func (tab *Tabulator) ToTable(records [][]string) string {
	tab.calculateFieldLengths(records)

	output := strings.Builder{}
	output.WriteString(tab.GetTableHeader() + "\n")

	if len(records) == 0 {
		output.WriteString("No records found\n")
	}

	// write the records to the string builder
	for _, record := range records {
		output.WriteString(tab.ToRow(record) + "\n")
	}

	return output.String()
}

func (tab *Tabulator) ToSegmentedTable(records [][]string, segmentIndex int) string {
	tab.calculateFieldLengths(records)

	output := strings.Builder{}
	output.WriteString(tab.GetTableHeader() + "\n")

	if len(records) == 0 {
		output.WriteString("No records found\n")
	}

	segmentValue := ""
	isValidSegmentation := false
	if segmentIndex >= 0 && segmentIndex < len(records[0]) {
		segmentValue = records[0][segmentIndex]
		isValidSegmentation = true
	}

	for _, record := range records {
		if isValidSegmentation && record[segmentIndex] != segmentValue {
			output.WriteString(tab.GetTableHeader() + "\n")
			segmentValue = record[segmentIndex]
		}

		output.WriteString(tab.ToRow(record) + "\n")
	}

	return output.String()
}

func (tab *Tabulator) calculateFieldLengths(records [][]string) {
	for _, fields := range records {
		for i := 0; i < len(fields); i++ {
			if len(fields[i]) > tab.columnDefinitions[i].fieldLength {
				tab.columnDefinitions[i].fieldLength = len(fields[i])
			}
		}
	}
}
