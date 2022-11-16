package components

import (
	. "Goreact"
	"goreact_qt/internal/native_comp"
)

var Window = FCC("Qt Window", func(children ...Node) Node {
	return NativeEl(native_comp.WindowName, ChildrenProps{
		Children: children,
	})
})
