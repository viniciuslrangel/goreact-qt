package components

import (
	. "github.com/viniciuslrangel/goreact"
	"github.com/viniciuslrangel/goreact-qt/internal/native_comp"
)

type LabelProps = native_comp.LabelProps

var Label = FCP("Qt Label", func(props LabelProps) Node {
	return NativeEl(native_comp.LabelName, props)
})
