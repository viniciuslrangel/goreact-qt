package components

import (
	. "github.com/viniciuslrangel/goreact"
	"github.com/viniciuslrangel/goreact-qt/internal/native_comp"
)

type ButtonProps = native_comp.ButtonProps

var Button = FCP("Qt Button", func(props ButtonProps) Node {
	return NativeEl(native_comp.ButtonName, props)
})
