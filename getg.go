package threadMeta

import (
	"reflect"
	"unsafe"
)

func g_ptr() unsafe.Pointer

var (
	g_type          reflect.Type
	g_goid_offset   uintptr
	g_labels_offset uintptr
)

func init() {
	for _, ts := range typesByString("*runtime.g") {
		typ := reflect.TypeOf(0)
		rt := (*iface)(unsafe.Pointer(&typ)) // reflect.Type
		rt.data = ts
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		goidF, ok := typ.FieldByName("goid")
		if !ok {
			continue
		}
		labelsF, ok := typ.FieldByName("labels")
		if !ok {
			continue
		}
		g_type = typ
		g_goid_offset = goidF.Offset
		g_labels_offset = labelsF.Offset
		break
	}
	if g_type == nil {
		panic("type·runtime·g fetch failed")
	}
}

type g struct {
	goid   uint64
	labels *unsafe.Pointer // profiler labels
}

func getg() g {
	ptr := g_ptr()
	if ptr == nil {
		panic("get g pointer fail")
	}
	return g{
		goid:   *(*uint64)(unsafe.Pointer(uintptr(ptr) + g_goid_offset)),
		labels: (*unsafe.Pointer)(unsafe.Pointer(uintptr(ptr) + g_labels_offset)),
	}
}

//go:norace
func (gp g) getLabels() unsafe.Pointer {
	return *gp.labels
}

//go:norace
func (gp g) setLabels(labels unsafe.Pointer) {
	*gp.labels = labels
}
