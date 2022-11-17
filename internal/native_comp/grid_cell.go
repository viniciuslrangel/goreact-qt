package native_comp

import (
	"github.com/therecipe/qt/widgets"
	. "github.com/viniciuslrangel/goreact"
	"log"
)

type gridCell struct {
	render QtRender

	props GridCellProps

	parent *widgets.QGridLayout
	child  widgets.QWidget_ITF
}

type GridCellProps struct {
	ChildrenProps
	Row        int
	Column     int
	RowSpan    int
	ColumnSpan int
}

const GridCellName = "qt_grid_cell"

func init() {
	AllComponents[GridCellName] = func(render QtRender) Widget {
		return &gridCell{
			render: render,
		}
	}
}

func (g *gridCell) GetName() string {
	return GridCellName
}

func (g *gridCell) UpdateElement(node *NodeData) {
	nProps := node.Props.(GridCellProps)
	changed := false
	if nProps.RowSpan <= 0 {
		nProps.RowSpan = 1
	}
	if g.props.ColumnSpan <= 0 {
		g.props.ColumnSpan = 1
	}

	if nProps.RowSpan != g.props.RowSpan {
		changed = true
	} else if nProps.ColumnSpan != g.props.ColumnSpan {
		changed = true
	} else if nProps.Row != g.props.Row {
		changed = true
	} else if nProps.Column != g.props.Column {
		changed = true
	}
	g.props = nProps
	if changed && g.child != nil {
		g.parent.RemoveWidget(g.child)
		g.insertElement()
	}
}

func (g *gridCell) OnWidgetCreated(node *NodeData) {
	if node.Parent.NativeTyp != GridLayoutName {
		log.Printf("GridCell requires a GridLayout as parent\n")
		return
	}
	g.parent = node.Parent.NativeElement.(*gridLayout).layout
}

func (g *gridCell) OnWidgetRemoved(node *NodeData) {
	if g.child != nil {
		g.parent.RemoveWidget(g.child)
	}
	g.parent = nil
}

func (g *gridCell) AddQtWidget(child widgets.QWidget_ITF) {
	if g.child != nil {
		log.Printf("GridCell already has a child\n")
		return
	}
	g.child = child
	g.insertElement()
}

func (g *gridCell) RemoveQtWidget(child widgets.QWidget_ITF) {
	if g.child != child {
		log.Printf("GridCell does not have this child\n")
		return
	}
	g.parent.RemoveWidget(g.child)
	g.child = nil
}

func (g *gridCell) insertElement() {
	g.parent.AddWidget3(g.child, g.props.Row, g.props.Column, g.props.RowSpan, g.props.ColumnSpan, 0)
}
