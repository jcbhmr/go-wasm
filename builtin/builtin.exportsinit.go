package builtin

import (
	"runtime/cgo"

	"go.bytecodealliance.org/cm"
)

func init() {
	Exports.Error.Destructor = func(self cm.Rep) {
		cgo.Handle(self).Delete()
	}
	Exports.Error.Error2 = func(self cm.Rep) (result string) {
		return cgo.Handle(self).Value().(error).Error()
	}
}
