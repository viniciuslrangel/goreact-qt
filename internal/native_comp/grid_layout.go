package native_comp

import (
	"github.com/therecipe/qt/widgets"
	. "github.com/viniciuslrangel/goreact"
	"log"
)

type gridLayout struct {
	render QtRender
	widget *widgets.QWidget
	layout *widgets.QGridLayout
	parent ParentWidget
}

const GridLayoutName = "qt_grid_layout"

func init() {
	AllComponents[GridLayoutName] = func(render QtRender) Widget {
		return &gridLayout{
			render: render,
		}
	}
}

func (g *gridLayout) GetName() string {
	return GridLayoutName
}

func (g *gridLayout) UpdateElement(_ *NodeData) {
}

func (g *gridLayout) OnWidgetCreated(node *NodeData) {
	g.widget = widgets.NewQWidget(nil, 0)
	g.layout = widgets.NewQGridLayout(g.widget)
	g.widget.SetLayout(g.layout)
	g.parent = g.render.FindFirstContainer(node)
	if g.parent == nil {
		log.Printf("GridLayout requires a container node\n")
		return
	}
	g.parent.AddQtWidget(g.widget)
}

func (g *gridLayout) OnWidgetRemoved(node *NodeData) {
	g.widget.DeleteLater()
	if g.parent == nil {
		return
	}
	g.parent.RemoveQtWidget(g.widget)
}

func (g *gridLayout) AddQtWidget(widget widgets.QWidget_ITF, row, column, rowSpan, colSpawn int) {
	g.layout.AddWidget3(widget, row, column, rowSpan, colSpawn, 0)
}

func (g *gridLayout) RemoveWidget(child widgets.QWidget_ITF, row, column int) {
	g.layout.RemoveWidget(child)
	item := g.layout.ItemAtPosition(row, column)
	g.layout.RemoveItem(item)
}
