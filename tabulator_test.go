package tabulator

import (
	"os"
	"testing"
)

type TestData struct {
	symbol   string
	date     int
	volume   int
	avgPrice float64
	returns  float64
	rate     float64
}

func TestTableOutput(test *testing.T) {
	tab := NewTabulator("Column 1", "Column 2", "Column 3")

	records := [][]string{
		{"Field 1", "Field 2", "Field 3"},
		{"Field 1", "Field 2", "Field 3"},
		{"Field 1", "Field 2", "Field 3"},
	}

	for i, r := range records {
		tab.AddRow()
		tab.AddValueByColumnIndex(i, 0, r[0])
		tab.AddValueByColumnIndex(i, 1, r[1])
		tab.AddValueByColumnIndex(i, 2, r[2])
	}

	output := tab.ToTable()
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
	tab := NewTabulator()
	tab.AddColumn(NewStringColumn("Symbol"))
	tab.AddColumn(NewIntColumn("Date"))
	tab.AddColumn(NewIntColumn("Volume"))
	tab.AddColumn(NewCurrencyColumn("Avg Price"))
	tab.AddColumn(NewPercentageColumn("Returns"))
	tab.AddColumn(NewFloatColumn("Rate"))

	records := []TestData{
		{"ASDF", 20200930, 123456789, 1234.56, 0.082, 33.19},
		{"QWER", 20200930, 987654, 10112.50, -0.001, 10.15},
		{"ASDF", 20201001, 11112223, 112.31, 0.013, 99.43},
		{"QWER", 20201001, 231284, 891.17, 0.0031, 47.47},
	}

	for i, r := range records {
		tab.AddRow()
		tab.AddValueByColumnName(i, "Symbol", r.symbol)
		tab.AddValueByColumnName(i, "Date", r.date)
		tab.AddValueByColumnName(i, "Volume", r.volume)
		tab.AddValueByColumnName(i, "Avg Price", r.avgPrice)
		tab.AddValueByColumnName(i, "Returns", r.returns)
		tab.AddValueByColumnName(i, "Rate", r.rate)
	}

	output := tab.ToSegmentedTable(1)
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

func TestFixedColumnOutput(test *testing.T) {
	tab := NewTabulator()
	tab.AddColumn(NewFixedStringColumn("Fixed String", 16))
	tab.AddColumn(NewFixedIntColumn("Fixed Int", 12))
	tab.AddColumn(NewFixedFloatColumn("Fixed Float", 15))
	tab.AddColumn(NewFixedCurrencyColumn("Fixed Currency", 17))
	tab.AddColumn(NewFixedPercentageColumn("Fixed Percentage", 24))

	tab.AddRow()
	tab.AddValueByColumnIndex(0, 0, "Test Value 1")
	tab.AddValueByColumnIndex(0, 1, 123)
	tab.AddValueByColumnIndex(0, 2, 1.23)
	tab.AddValueByColumnIndex(0, 3, 12.11)
	tab.AddValueByColumnIndex(0, 4, 0.123)

	tab.AddRow()
	tab.AddValueByColumnIndex(1, 0, "Test Value 2")
	tab.AddValueByColumnIndex(1, 1, 456)
	tab.AddValueByColumnIndex(1, 2, 4.56)
	tab.AddValueByColumnIndex(1, 3, 34.33)
	tab.AddValueByColumnIndex(1, 4, 0.456)

	output := tab.ToTable()
	file, err := os.ReadFile("table_fixed_col_results.txt")
	if err != nil {
		test.Error("Unable to open table fixed column results file")
		return
	}

	expected := string(file)
	if output != expected {
		test.Errorf("Unexpected table output: %s", output)
	}
}
