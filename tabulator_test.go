package tabulator

import (
	"os"
	"testing"
)

func TestTableOutput(test *testing.T) {
	col1 := NewColumn("Column 1", 6, 8)
	col2 := NewColumn("Column 2", 6, 8)
	col3 := NewColumn("Column 3", 6, 8)

	tab := NewTabulator()
	tab.AddColumnDefinition(col1)
	tab.AddColumnDefinition(col2)
	tab.AddColumnDefinition(col3)

	records := [][]string{
		{"Field 1", "Field 2", "Field 3"},
		{"Field 1", "Field 2", "Field 3"},
		{"Field 1", "Field 2", "Field 3"},
	}

	output := tab.ToTable(records)
	file, err := os.ReadFile("table_results.txt")
	if err != nil {
		test.Errorf("Unable to open table results file")
		return
	}

	expected := string(file)
	if output != expected {
		test.Errorf("Unexpected table output: %s", output)
	}
}

func TestSegmentedTableOutput(test *testing.T) {
	tab := NewTabulator("Symbol", "Date", "Volume", "Avg Price", "Rate")
	records := [][]string{
		{"ASDF", "20200930", "123456789", "$1234.12", "33.19%"},
		{"QWER", "20200930", "987654", "$10112.50", "10.15%"},
		{"ASDF", "20201001", "11112223", "$112.31", "99.43%"},
		{"QWER", "20201001", "231284", "$891.17", "47.47%"},
	}

	output := tab.ToSegmentedTable(records, 1)
	file, err := os.ReadFile("table_segment_results.txt")
	if err != nil {
		test.Error("Unable to open table segment results file")
		return
	}

	expected := string(file)
	if output != expected {
		test.Errorf("Unexpected table output: %s", output)
	}
}
