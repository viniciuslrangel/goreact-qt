package components

import (
	. "github.com/viniciuslrangel/goreact"
	"github.com/viniciuslrangel/goreact-qt/internal/native_comp"
)

type cellProps = native_comp.GridCellProps

var _gridLayout = FCC("Qt Grid Layout", func(children ...Node) Node {
	return NativeEl(native_comp.GridLayoutName, ChildrenProps{
		Children: children,
	})
})

type gridLayoutCtor struct{}

var GridLayout = gridLayoutCtor{}

func (gridLayoutCtor) Keyed(key Key, children ...cellProps) Node {
	var cellList = make([]Node, len(children))
	for i, child := range children {
		cellList[i] = NativeEl(native_comp.GridCellName, child)
	}
	return _gridLayout.Keyed(key, cellList...)
}

func (gridLayoutCtor) New(children ...cellProps) Node {
	return GridLayout.Keyed(NoKey, children...)
}

//goland:noinspection GoExportedFuncWithUnexportedType
func GridCell(row, col int, child Node) cellProps {
	return native_comp.GridCellProps{
		Row:    row,
		Column: col,
		ChildrenProps: ChildrenProps{
			Children: []Node{child},
		},
	}
}
