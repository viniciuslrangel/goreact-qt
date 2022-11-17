package native_comp

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	. "github.com/viniciuslrangel/goreact"
	"log"
)

type label struct {
	render QtRender
	widget *widgets.QLabel

	props LabelProps
}

type LabelProps struct {
	Text string
}

const LabelName = "qt_label"

func init() {
	AllComponents[LabelName] = func(render QtRender) Widget {
		return &label{
			render: render,
		}
	}
}

func (l *label) GetName() string {
	return LabelName
}

func (l *label) UpdateElement(node *NodeData) {
	last := l.props
	props := node.Props.(LabelProps)
	l.props = props
	if last.Text != props.Text {
		l.widget.SetText(props.Text)
	}
}

func (l *label) OnWidgetCreated(node *NodeData) {
	l.widget = widgets.NewQLabel(nil, core.Qt__Widget)
	parent := l.render.FindFirstContainer(node)
	if parent == nil {
		log.Printf("Label requires a container node\n")
		return
	}
	parent.AddQtWidget(l.widget)
}

func (l *label) OnWidgetRemoved(node *NodeData) {
	parent := l.widget.ParentWidget()
	if parent == nil {
		log.Printf("Label has no parent\n")
		return
	}
	parent.Layout().RemoveWidget(l.widget)
}
