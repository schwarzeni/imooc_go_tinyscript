package util

// PriorityTable 优先级表
type PriorityTable struct {
	table []map[string]struct{}
}

// Size 返回 PriorityTable 大小
func (pt *PriorityTable) Size() int {
	return len(pt.table)
}

// GetByLevel 获取当前优先级的所有符号
func (pt *PriorityTable) GetByLevel(level int) map[string]struct{} {
	return pt.table[level]
}

// Has 判断符号 opt 在当前优先级中是否存在
func (pt *PriorityTable) Has(level int, opt string) bool {
	_, ok := pt.GetByLevel(level)[opt]
	return ok
}

func NewPrioriryTable() *PriorityTable {
	return &PriorityTable{
		table: []map[string]struct{}{
			{"&": {}, "|": {}, "^": {}},
			{"==": {}, "!=": {}, ">": {}, "<": {}, ">=": {}, "<=": {}},
			{"+": {}, "-": {}},
			{"*": {}, "/": {}},
			{"<<": {}, ">>": {}},
		},
	}
}
