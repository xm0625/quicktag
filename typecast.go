package quicktag

import (
	"reflect"
	"unsafe"
)

type emptyInterface struct {
	pt unsafe.Pointer
	pv unsafe.Pointer
}

func PointerOfType(t reflect.Type) unsafe.Pointer {
	p := *(*emptyInterface)(unsafe.Pointer(&t))
	return p.pv
}

func TypeCast(src interface{}, dstType reflect.Type) (dst interface{}) {
	srcType := reflect.TypeOf(src)
	eface := *(*emptyInterface)(unsafe.Pointer(&src))
	if srcType.Kind() == reflect.Ptr {
		eface.pt = PointerOfType(reflect.PtrTo(dstType))
	} else {
		eface.pt = PointerOfType(dstType)
	}
	dst = *(*interface{})(unsafe.Pointer(&eface))
	return
}
