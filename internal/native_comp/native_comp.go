package native_comp

import (
	"github.com/therecipe/qt/widgets"
	. "github.com/viniciuslrangel/goreact"
)

type Widget interface {
	GetName() string

	UpdateElement(node *NodeData)

	OnWidgetCreated(node *NodeData)
	OnWidgetRemoved(node *NodeData)
}

type ParentWidget interface {
	AddQtWidget(child widgets.QWidget_ITF)
	RemoveQtWidget(child widgets.QWidget_ITF)
}

type QtRender interface {
	FindFirstContainer(node *NodeData) ParentWidget
}

var AllComponents = make(map[string]func(render QtRender) Widget)
