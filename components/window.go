package components

import (
	. "github.com/viniciuslrangel/goreact"
	"github.com/viniciuslrangel/goreact-qt/internal/native_comp"
)

var Window = FCSC("Qt Window", func(child Node) Node {
	return NativeEl(native_comp.WindowName, ChildrenProps{
		Children: []Node{child},
	})
})
