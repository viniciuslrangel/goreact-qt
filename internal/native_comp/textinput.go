package native_comp

import (
	"github.com/therecipe/qt/widgets"
	"github.com/viniciuslrangel/goreact"
	. "github.com/viniciuslrangel/goreact/util"
	"log"
)

type textinput struct {
	render QtRender
	widget *widgets.QLineEdit
	parent ParentWidget

	props TextInputProps
}

type TextInputProps struct {
	Text          Optional[string]
	IsReadOnly    bool
	OnTextChanged func(string)
}

const TextInputName = "qt_textinput"

func init() {
	AllComponents[TextInputName] = func(render QtRender) Widget {
		return &textinput{
			render: render,
		}
	}
}

func (t *textinput) GetName() string {
	return TextInputName
}

func (t *textinput) UpdateElement(node *goreact.NodeData) {
	last := t.props
	props := node.Props.(TextInputProps)
	t.props = props
	if props.Text.IsPresent() {
		if !last.Text.IsPresent() || last.Text.Get() != props.Text.Get() {
			t.widget.SetText(props.Text.Get())
		}
	}
	if last.IsReadOnly != props.IsReadOnly {
		t.widget.SetReadOnly(props.IsReadOnly)
	}
}

func (t *textinput) OnWidgetCreated(node *goreact.NodeData) {
	t.widget = widgets.NewQLineEdit(nil)
	t.parent = t.render.FindFirstContainer(node)
	if t.parent == nil {
		log.Printf("TextInput requires a container node\n")
		return
	}
	t.parent.AddQtWidget(t.widget)

	t.widget.ConnectTextChanged(func(text string) {
		if t.props.OnTextChanged != nil {
			t.props.OnTextChanged(text)
		}
	})
}

func (t *textinput) OnWidgetRemoved(node *goreact.NodeData) {
	t.widget.DeleteLater()
	if t.parent == nil {
		return
	}
	t.parent.RemoveQtWidget(t.widget)
}
