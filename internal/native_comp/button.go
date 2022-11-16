package native_comp

import (
	. "Goreact"
	"github.com/therecipe/qt/widgets"
	"log"
)

type button struct {
	render QtRender
	widget *widgets.QPushButton
}

type ButtonProps struct {
	Label string
	Flat  bool
}

const ButtonName = "qt_button"

func init() {
	AllComponents[ButtonName] = func(render QtRender) Widget {
		return &button{
			render: render,
		}
	}
}

func (b *button) GetName() string {
	return ButtonName
}

func (b *button) UpdateElement(element *NodeData) {
	props := element.Props.(ButtonProps)
	b.widget.SetText(props.Label)
}

func (b *button) OnWidgetCreated(node *NodeData) {
	b.widget = widgets.NewQPushButton(nil)
	parent := b.render.FindFirstContainer(node)
	if parent == nil {
		log.Printf("Button requires a container node\n")
		return
	}
	parent.AddQtWidget(b.widget)
}

func (b *button) OnWidgetRemoved(node *NodeData) {
	parent := b.widget.ParentWidget()
	if parent == nil {
		log.Printf("Button has no parent\n")
		return
	}
	parent.Layout().RemoveWidget(b.widget)
}
