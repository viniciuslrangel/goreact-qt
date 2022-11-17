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
	parent := g.render.FindFirstContainer(node)
	if parent == nil {
		log.Printf("GridLayout requires a container node\n")
		return
	}
	parent.AddQtWidget(g.widget)
}

func (g *gridLayout) OnWidgetRemoved(node *NodeData) {
	parent := g.widget.ParentWidget()
	if parent == nil {
		log.Printf("GridLayout has no parent\n")
		return
	}
	parent.Layout().RemoveWidget(g.widget)
}
