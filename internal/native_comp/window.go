package native_comp

import (
	"github.com/therecipe/qt/widgets"
	. "github.com/viniciuslrangel/goreact"
)

/**

The window component is a bit special, because it is the root, so it doesn't have a parent and
we must use the child parameter instead of parent

*/

type window struct {
	render QtRender
	widget *widgets.QMainWindow
}

const WindowName = "qt_window"

func init() {
	AllComponents[WindowName] = func(render QtRender) Widget {
		return &window{
			render: render,
		}
	}
}

func (w *window) GetName() string {
	return WindowName
}

func (w *window) UpdateElement(node *NodeData) {
	//TODO implement me
}

func (w *window) OnWidgetCreated(node *NodeData) {
	w.widget = widgets.NewQMainWindow(nil, 0)
	w.widget.Show()
}

func (w *window) OnWidgetRemoved(node *NodeData) {
	w.widget.Close()
	w.widget.DeleteLater()
	w.widget = nil
}

func (w *window) AddQtWidget(child widgets.QWidget_ITF) {
	w.widget.SetCentralWidget(child)
}

func (w *window) RemoveQtWidget(child widgets.QWidget_ITF) {
	w.widget.Layout().RemoveWidget(child)
}
