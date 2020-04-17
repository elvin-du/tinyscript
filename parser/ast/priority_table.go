package ast

var PriorityTable = NewPriorityTable()

type priorityTable struct {
	table [][]string
}

func NewPriorityTable() *priorityTable {
	return &priorityTable{[][]string{
		[]string{"&", "|", "^"},
		[]string{"==", "!=", ">", "<", ">=", "<="},
		[]string{"+", "-"},
		[]string{"*", "/"},
		[]string{"<<", ">>"},
	}}
}

func (pt *priorityTable) Size() int {
	return len(pt.table)
}
func (pt *priorityTable) Get(level int) []string {
	return pt.table[level]
}
func (pt *priorityTable) IsContain(level int, key string) bool {
	strs := pt.Get(level)
	for _, str := range strs {
		if str == key {
			return true
		}
	}

	return false
}
