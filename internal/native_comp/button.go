package native_comp

import (
	"github.com/therecipe/qt/widgets"
	. "github.com/viniciuslrangel/goreact"
	"log"
)

type button struct {
	render QtRender
	widget *widgets.QPushButton
	parent ParentWidget

	props ButtonProps
}

type ButtonProps struct {
	Label     string
	OnClicked func()
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

func (b *button) UpdateElement(node *NodeData) {
	last := b.props
	props := node.Props.(ButtonProps)
	b.props = props
	if last.Label != props.Label {
		b.widget.SetText(props.Label)
	}
}

func (b *button) OnWidgetCreated(node *NodeData) {
	b.widget = widgets.NewQPushButton(nil)
	b.parent = b.render.FindFirstContainer(node)
	if b.parent == nil {
		log.Printf("Button requires a container node\n")
		return
	}
	b.parent.AddQtWidget(b.widget)

	b.widget.ConnectClicked(func(_ bool) {
		if b.props.OnClicked != nil {
			b.props.OnClicked()
		}
	})
}

func (b *button) OnWidgetRemoved(node *NodeData) {
	b.widget.DeleteLater()
	if b.parent == nil {
		return
	}
	b.parent.RemoveQtWidget(b.widget)
}
