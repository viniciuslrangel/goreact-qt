package components

import (
	. "github.com/viniciuslrangel/goreact"
	"github.com/viniciuslrangel/goreact-qt/internal/native_comp"
)

type TextInputProps = native_comp.TextInputProps

var TextInput = FCP("Qt TextInput", func(props TextInputProps) Node {
	return NativeEl(native_comp.TextInputName, props)
})
