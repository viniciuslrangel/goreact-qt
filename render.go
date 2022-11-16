package goreact_qt

import (
	. "Goreact"
	"github.com/therecipe/qt/widgets"
	"goreact_qt/internal/native_comp"
	"os"
)

type qtRender struct {
	engine *Engine
}

func (q *qtRender) FindFirstContainer(e *NodeData) native_comp.ParentWidget {
	if e == nil {
		return nil
	}
	containerWidget, ok := e.NativeElement.(native_comp.ParentWidget)
	if ok {
		return containerWidget
	}
	return q.FindFirstContainer(e.Parent)
}

func (q *qtRender) SetEngine(e *Engine) {
	q.engine = e
}

func (q *qtRender) InsertNode(node *NodeData) {
	comp := native_comp.AllComponents[node.NativeTyp]
	nativeEl := comp(q)
	node.NativeElement = nativeEl
	nativeEl.OnWidgetCreated(node)
}

func (q *qtRender) RemoveNode(node *NodeData) {
	widget := node.NativeElement.(native_comp.Widget)
	widget.OnWidgetRemoved(node)
}

func (q *qtRender) UpdateElement(node *NodeData) {
	widget := node.NativeElement.(native_comp.Widget)
	widget.UpdateElement(node)

	cProps, hasChildren := node.Props.(ChildrenProps)
	if hasChildren {
		node.Children = cProps.Children
	}
}

func Render(root Node) error {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	var engine qtRender
	NewEngine(&engine).Render(root)

	app.Exec()

	return nil
}
