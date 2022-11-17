package native_comp

import (
	"github.com/therecipe/qt/widgets"
	. "github.com/viniciuslrangel/goreact"
	"log"
)

type button struct {
	render QtRender
	widget *widgets.QPushButton

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

func (b *button) UpdateElement(element *NodeData) {
	last := b.props
	props := element.Props.(ButtonProps)
	b.props = props
	if last.Label != props.Label {
		b.widget.SetText(props.Label)
	}
}

func (b *button) OnWidgetCreated(node *NodeData) {
	b.widget = widgets.NewQPushButton(nil)
	parent := b.render.FindFirstContainer(node)
	if parent == nil {
		log.Printf("Button requires a container node\n")
		return
	}
	parent.AddQtWidget(b.widget)

	b.widget.ConnectClicked(func(_ bool) {
		if b.props.OnClicked != nil {
			b.props.OnClicked()
		}
	})
}

func (b *button) OnWidgetRemoved(node *NodeData) {
	parent := b.widget.ParentWidget()
	if parent == nil {
		log.Printf("Button has no parent\n")
		return
	}
	parent.Layout().RemoveWidget(b.widget)
}
