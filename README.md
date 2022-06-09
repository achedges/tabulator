# tabulator
`tabulator` is a simple library that allows you to output record sets as a column-justified table.

As an example, take the following data set (already string-formatted):
```
records := [][]string{
    {"ASDF", "20200930", "123456789", "$1234.12", "33.19%"},
    {"QWER", "20200930", "987654", "$10112.50", "10.15%"},
    {"ASDF", "20201001", "11112223", "$112.31", "99.43%"},
    {"QWER", "20201001", "231284", "$891.17", "47.47%"},
}
```

## Creating a basic table
```
tab := NewTabulator("Symbol", "Date", "Volume", "Avg Price", "Rate")
output := tab.ToTable(records)
```
```
Symbol    Date        Volume       Avg Price    Rate      
----------------------------------------------------------
ASDF      20200930    123456789    $1234.12     33.19%    
QWER      20200930    987654       $10112.50    10.15%    
ASDF      20201001    11112223     $112.31      99.43%    
QWER      20201001    231284       $891.17      47.47%
```
## Creating a segmented table
A segmented table is one where the table header is repeated whenever a change is detected in the value of the field at the `segmentIndex`.  The following example shows a table segmented by the `Date` field.
```
tab := NewTabulator("Symbol", "Date", "Volume", "Avg Price", "Rate")
output := tab.ToSegmentedTable(records, 1)
```
```
Symbol    Date        Volume       Avg Price    Rate      
----------------------------------------------------------
ASDF      20200930    123456789    $1234.12     33.19%    
QWER      20200930    987654       $10112.50    10.15%    

Symbol    Date        Volume       Avg Price    Rate      
----------------------------------------------------------
ASDF      20201001    11112223     $112.31      99.43%    
QWER      20201001    231284       $891.17      47.47%    
```
