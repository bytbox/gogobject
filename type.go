// Converting GIR type to Go/C representation

package main

import (
	"bytes"
	"fmt"
	"strings"
	"unsafe"
)

func ForcePointer(x string) string {
	if x == "unsafe.Pointer" {
		return x
	}
	if !strings.HasPrefix(x, "*") {
		return "*" + x
	}
	return x
}

type TypeFlags int

const (
	TypeNone    TypeFlags = 0
	TypePointer TypeFlags = 1 << iota
	TypeReturn
	TypeListMember
	TypeReceiver
	TypeExact
)

//------------------------------------------------------------------
// Cgo Type (C type in Go)
//------------------------------------------------------------------

func CgoType(ti *TypeInfo, flags TypeFlags) string {
	var out bytes.Buffer

	switch tag := ti.Tag(); tag {
	case TYPE_TAG_VOID:
		if ti.IsPointer() {
			out.WriteString("unsafe.Pointer")
			break
		}
		panic("Non-pointer void type is not supported in cgo")
	case TYPE_TAG_UTF8, TYPE_TAG_FILENAME:
		out.WriteString("*C.char")
	case TYPE_TAG_ARRAY:
		switch ti.ArrayType() {
		case ARRAY_TYPE_C:
			out.WriteString("*")
			out.WriteString(CgoType(ti.ParamType(0), flags))
		case ARRAY_TYPE_ARRAY:
			out.WriteString("*C.GArray")
		case ARRAY_TYPE_PTR_ARRAY:
			out.WriteString("*C.GPtrArray")
		case ARRAY_TYPE_BYTE_ARRAY:
			out.WriteString("*C.GByteArray")
		}
	case TYPE_TAG_GLIST:
		out.WriteString("*C.GList")
	case TYPE_TAG_GSLIST:
		out.WriteString("*C.GSList")
	case TYPE_TAG_GHASH:
		out.WriteString("*C.GHashTable")
	case TYPE_TAG_ERROR:
		out.WriteString("*C.GError")
	case TYPE_TAG_INTERFACE:
		if ti.IsPointer() {
			flags |= TypePointer
		}
		out.WriteString(CgoTypeForInterface(ti.Interface(), flags))
	default:
		if ti.IsPointer() {
			out.WriteString("*")
		}
		out.WriteString(CgoTypeForTag(tag, flags))
	}

	return out.String()
}

func CgoTypeForInterface(bi *BaseInfo, flags TypeFlags) string {
	var out bytes.Buffer

	switch bi.Type() {
	case INFO_TYPE_CALLBACK:
		out.WriteString("unsafe.Pointer")
	default:
		ns := bi.Namespace()
		nm := bi.Name()
		fullnm := strings.ToLower(ns) + "." + nm

		_, disguised := GConfig.Sys.DisguisedTypes[fullnm]
		if flags&TypePointer != 0 && !disguised {
			out.WriteString("*")
		}

		out.WriteString("C.")
		out.WriteString(DefaultRepository().CPrefix(ns))
		out.WriteString(bi.Name())
	}
	return out.String()
}

func CgoTypeForTag(tag TypeTag, flags TypeFlags) string {
	var out bytes.Buffer
	p := PrinterTo(&out)

	if flags & TypePointer != 0 {
		p("*")
	}

	switch tag {
	case TYPE_TAG_BOOLEAN:
		p("C.int")
	case TYPE_TAG_INT8:
		p("C.int8_t")
	case TYPE_TAG_UINT8:
		p("C.uint8_t")
	case TYPE_TAG_INT16:
		p("C.int16_t")
	case TYPE_TAG_UINT16:
		p("C.uint16_t")
	case TYPE_TAG_INT32:
		p("C.int32_t")
	case TYPE_TAG_UINT32:
		p("C.uint32_t")
	case TYPE_TAG_INT64:
		p("C.int64_t")
	case TYPE_TAG_UINT64:
		p("C.uint64_t")
	case TYPE_TAG_FLOAT:
		p("C.float")
	case TYPE_TAG_DOUBLE:
		p("C.double")
	case TYPE_TAG_GTYPE:
		p("C.GType")
	case TYPE_TAG_UNICHAR:
		p("C.uint32_t")
	default:
		panic("unreachable")
	}

	return out.String()
}

//------------------------------------------------------------------
// C Type
//------------------------------------------------------------------

func CTypeForTag(tag TypeTag, flags TypeFlags) string {
	return CgoTypeForTag(tag, flags)[2:]
}

func CType(ti *TypeInfo, flags TypeFlags) string {
	var out bytes.Buffer

	switch tag := ti.Tag(); tag {
	case TYPE_TAG_VOID:
		if ti.IsPointer() {
			out.WriteString("void*")
			break
		}
		out.WriteString("void")
	case TYPE_TAG_UTF8, TYPE_TAG_FILENAME:
		out.WriteString("char*")
	case TYPE_TAG_ARRAY:
		switch ti.ArrayType() {
		case ARRAY_TYPE_C:
			out.WriteString(CType(ti.ParamType(0), flags))
			out.WriteString("*")
		case ARRAY_TYPE_ARRAY:
			out.WriteString("GArray*")
		case ARRAY_TYPE_PTR_ARRAY:
			out.WriteString("GPtrArray*")
		case ARRAY_TYPE_BYTE_ARRAY:
			out.WriteString("GByteArray*")
		}
	case TYPE_TAG_GLIST:
		out.WriteString("GList*")
	case TYPE_TAG_GSLIST:
		out.WriteString("GSList*")
	case TYPE_TAG_GHASH:
		out.WriteString("GHashTable*")
	case TYPE_TAG_ERROR:
		out.WriteString("GError*")
	case TYPE_TAG_INTERFACE:
		if ti.IsPointer() {
			flags |= TypePointer
		}
		out.WriteString(CTypeForInterface(ti.Interface(), flags))
	default:
		out.WriteString(CTypeForTag(tag, flags))
		if ti.IsPointer() {
			out.WriteString("*")
		}
	}

	return out.String()
}

func CTypeForInterface(bi *BaseInfo, flags TypeFlags) string {
	var out bytes.Buffer

	ns := bi.Namespace()
	nm := bi.Name()
	fullnm := strings.ToLower(ns) + "." + nm
	out.WriteString(DefaultRepository().CPrefix(ns))
	out.WriteString(bi.Name())

	_, disguised := GConfig.Sys.DisguisedTypes[fullnm]
	if flags&TypePointer != 0 && !disguised {
		out.WriteString("*")
	}

	return out.String()
}

//------------------------------------------------------------------
// Go Type
//------------------------------------------------------------------

func GoTypeForInterface(bi *BaseInfo, flags TypeFlags) string {
	var out bytes.Buffer
	printf := PrinterTo(&out)
	ns := bi.Namespace()
	fullnm := strings.ToLower(ns) + "." + bi.Name()

	if flags&TypeListMember != 0 {
		switch bi.Type() {
		case INFO_TYPE_OBJECT, INFO_TYPE_INTERFACE:
			return GoTypeForInterface(bi, TypePointer|TypeReturn)
		default:
			return GoTypeForInterface(bi, TypeReturn)
		}
	}

	switch t := bi.Type(); t {
	case INFO_TYPE_OBJECT, INFO_TYPE_INTERFACE:
		if flags&TypeExact != 0 {
			// exact type for object/interface is always an unsafe.Pointer
			printf("unsafe.Pointer")
			break
		}

		if flags&(TypeReturn|TypeReceiver) != 0 && flags&TypePointer != 0 {
			// receivers and return values are actual types,
			// and a pointer most likely
			printf("*")
		}
		if ns != Config.Namespace {
			// prepend foreign types with appropriate namespace
			printf("%s.", strings.ToLower(ns))
		}
		printf(bi.Name())
		if flags&(TypeReturn|TypeReceiver) == 0 {
			// ordinary function arguments are substituted by their *Like
			// counterparts
			printf("Like")
		}
		if flags&TypeReceiver != 0 && t == INFO_TYPE_INTERFACE {
			// special case for interfaces, we use *Impl structures
			// as receivers
			printf("Impl")
		}
	case INFO_TYPE_CALLBACK:
		if flags&TypeExact != 0 {
			printf("unsafe.Pointer")
			break
		}
		goto handle_default
	case INFO_TYPE_STRUCT:
		if ns == "cairo" {
			printf(CairoGoTypeForInterface(bi, flags))
			break
		}
		goto handle_default
	default:
		goto handle_default
	}
	return out.String()
handle_default:
	_, disguised := GConfig.Sys.DisguisedTypes[fullnm]
	if flags&TypePointer != 0 && !disguised {
		printf("*")
	}
	if ns != Config.Namespace {
		printf("%s.", strings.ToLower(ns))
	}
	printf(bi.Name())
	return out.String()
}

func GoType(ti *TypeInfo, flags TypeFlags) string {
	var out bytes.Buffer

	switch tag := ti.Tag(); tag {
	case TYPE_TAG_VOID:
		if ti.IsPointer() {
			out.WriteString("unsafe.Pointer")
			break
		}
		panic("Non-pointer void type is not supported")
	case TYPE_TAG_UTF8, TYPE_TAG_FILENAME:
		if flags&TypeExact != 0 {
			out.WriteString("unsafe.Pointer")
		} else {
			out.WriteString("string")
		}
	case TYPE_TAG_ARRAY:
		size := ti.ArrayFixedSize()
		if size != -1 {
			fmt.Fprintf(&out, "[%d]", size)
		} else {
			if flags&TypeExact != 0 {
				out.WriteString("unsafe.Pointer")
			} else {
				out.WriteString("[]")
			}
		}
		out.WriteString(GoType(ti.ParamType(0), flags))
	case TYPE_TAG_GLIST:
		if flags&TypeExact != 0 {
			out.WriteString("unsafe.Pointer")
		} else {
			out.WriteString("[]")
			out.WriteString(GoType(ti.ParamType(0), flags|TypeListMember))
		}
	case TYPE_TAG_GSLIST:
		if flags&TypeExact != 0 {
			out.WriteString("unsafe.Pointer")
		} else {
			out.WriteString("[]")
			out.WriteString(GoType(ti.ParamType(0), flags|TypeListMember))
		}
	case TYPE_TAG_GHASH:
		if flags&TypeExact != 0 {
			out.WriteString("unsafe.Pointer")
		} else {
			out.WriteString("map[")
			out.WriteString(GoType(ti.ParamType(0), flags))
			out.WriteString("]")
			out.WriteString(GoType(ti.ParamType(1), flags))
		}
	case TYPE_TAG_ERROR:
		// not used?
		out.WriteString("error")
	case TYPE_TAG_INTERFACE:
		if ti.IsPointer() {
			flags |= TypePointer
		}
		out.WriteString(GoTypeForInterface(ti.Interface(), flags))
	default:
		if ti.IsPointer() {
			flags |= TypePointer
		}
		out.WriteString(GoTypeForTag(tag, flags))
	}

	return out.String()
}

func GoTypeForTag(tag TypeTag, flags TypeFlags) string {
	var out bytes.Buffer
	p := PrinterTo(&out)

	if flags & TypePointer != 0 {
		p("*")
	}

	if flags&TypeExact != 0 {
		switch tag {
		case TYPE_TAG_BOOLEAN:
			p("int32") // sadly
		case TYPE_TAG_INT8:
			p("int8")
		case TYPE_TAG_UINT8:
			p("uint8")
		case TYPE_TAG_INT16:
			p("int16")
		case TYPE_TAG_UINT16:
			p("uint16")
		case TYPE_TAG_INT32:
			p("int32")
		case TYPE_TAG_UINT32:
			p("uint32")
		case TYPE_TAG_INT64:
			p("int64")
		case TYPE_TAG_UINT64:
			p("uint64")
		case TYPE_TAG_FLOAT:
			p("float32")
		case TYPE_TAG_DOUBLE:
			p("float64")
		case TYPE_TAG_GTYPE:
			if Config.Namespace != "GObject" {
				p("gobject.Type")
			} else {
				p("Type")
			}
		case TYPE_TAG_UNICHAR:
			p("rune")
		default:
			panic("unreachable")
		}
	} else {
		switch tag {
		case TYPE_TAG_BOOLEAN:
			p("bool")
		case TYPE_TAG_INT8:
			p("int")
		case TYPE_TAG_UINT8:
			p("int")
		case TYPE_TAG_INT16:
			p("int")
		case TYPE_TAG_UINT16:
			p("int")
		case TYPE_TAG_INT32:
			p("int")
		case TYPE_TAG_UINT32:
			p("int")
		case TYPE_TAG_INT64:
			p("int64")
		case TYPE_TAG_UINT64:
			p("uint64")
		case TYPE_TAG_FLOAT:
			p("float64")
		case TYPE_TAG_DOUBLE:
			p("float64")
		case TYPE_TAG_GTYPE:
			if Config.Namespace != "GObject" {
				p("gobject.Type")
			} else {
				p("Type")
			}
		case TYPE_TAG_UNICHAR:
			p("rune")
		default:
			panic("unreachable")
		}
	}

	return out.String()
}

//------------------------------------------------------------------
// Simple Cgo Type (for exported functions)
//------------------------------------------------------------------

func SimpleCgoType(ti *TypeInfo, flags TypeFlags) string {
	tag := ti.Tag()
	switch tag {
	case TYPE_TAG_VOID:
		if ti.IsPointer() {
			return "unsafe.Pointer"
		}
		panic("Non-pointer void type is not supported")
	case TYPE_TAG_INTERFACE:
		bi := ti.Interface()
		switch bi.Type() {
		case INFO_TYPE_ENUM, INFO_TYPE_FLAGS:
			ei := ToEnumInfo(bi)
			return GoTypeForTag(ei.StorageType(), flags | TypeExact)
		case INFO_TYPE_STRUCT:
			ns := bi.Namespace()
			nm := bi.Name()
			fullnm := strings.ToLower(ns) + "." + nm
			if _, ok := GConfig.Sys.DisguisedTypes[fullnm]; ok {
				return "unsafe.Pointer"
			}
		}
	}
	if !strings.HasPrefix(CgoType(ti, flags), "*") {
		return GoTypeForTag(tag, flags | TypeExact)
	}
	return "unsafe.Pointer"
}

//------------------------------------------------------------------
// Type sizes
//------------------------------------------------------------------

func TypeSizeForInterface(bi *BaseInfo, flags TypeFlags) int {
	ptrsize := int(unsafe.Sizeof(unsafe.Pointer(nil)))
	if flags&TypePointer != 0 {
		return ptrsize
	}

	switch t := bi.Type(); t {
	case INFO_TYPE_OBJECT, INFO_TYPE_INTERFACE:
		return ptrsize
	case INFO_TYPE_STRUCT:
		si := ToStructInfo(bi)
		return si.Size()
	case INFO_TYPE_UNION:
		ui := ToUnionInfo(bi)
		return ui.Size()
	case INFO_TYPE_ENUM, INFO_TYPE_FLAGS:
		ei := ToEnumInfo(bi)
		return TypeSizeForTag(ei.StorageType(), flags)
	case INFO_TYPE_CALLBACK:
		return ptrsize
	}
	panic("unreachable: " + bi.Type().String())
}

// returns the size of a type, works only for TypeExact
func TypeSize(ti *TypeInfo, flags TypeFlags) int {
	ptrsize := int(unsafe.Sizeof(unsafe.Pointer(nil)))
	switch tag := ti.Tag(); tag {
	case TYPE_TAG_VOID:
		if ti.IsPointer() {
			return ptrsize
		}
		panic("Non-pointer void type is not supported")
	case TYPE_TAG_UTF8, TYPE_TAG_FILENAME, TYPE_TAG_GLIST,
		TYPE_TAG_GSLIST, TYPE_TAG_GHASH:
		return ptrsize
	case TYPE_TAG_ARRAY:
		size := ti.ArrayFixedSize()
		if size != -1 {
			return size * TypeSize(ti.ParamType(0), flags)
		}
		return ptrsize
	case TYPE_TAG_INTERFACE:
		if ti.IsPointer() {
			flags |= TypePointer
		}
		return TypeSizeForInterface(ti.Interface(), flags)
	default:
		if ti.IsPointer() {
			flags |= TypePointer
		}
		return TypeSizeForTag(tag, flags)
	}
	panic("unreachable: " + ti.Tag().String())
}

func TypeSizeForTag(tag TypeTag, flags TypeFlags) int {
	ptrsize := int(unsafe.Sizeof(unsafe.Pointer(nil)))
	if flags&TypePointer != 0 {
		return ptrsize
	}

	switch tag {
	case TYPE_TAG_BOOLEAN:
		return 4
	case TYPE_TAG_INT8:
		return 1
	case TYPE_TAG_UINT8:
		return 1
	case TYPE_TAG_INT16:
		return 2
	case TYPE_TAG_UINT16:
		return 2
	case TYPE_TAG_INT32:
		return 4
	case TYPE_TAG_UINT32:
		return 4
	case TYPE_TAG_INT64:
		return 8
	case TYPE_TAG_UINT64:
		return 8
	case TYPE_TAG_FLOAT:
		return 4
	case TYPE_TAG_DOUBLE:
		return 8
	case TYPE_TAG_GTYPE:
		return ptrsize
	case TYPE_TAG_UNICHAR:
		return 4
	}
	panic("unreachable: " + tag.String())
}

//------------------------------------------------------------------
// Type needs wrapper?
//------------------------------------------------------------------

func TypeNeedsWrapper(ti *TypeInfo) bool {
	switch tag := ti.Tag(); tag {
	case TYPE_TAG_VOID:
		if ti.IsPointer() {
			return false
		}
		panic("Non-pointer void type is not supported")
	case TYPE_TAG_UTF8, TYPE_TAG_FILENAME, TYPE_TAG_GLIST,
		TYPE_TAG_GSLIST, TYPE_TAG_GHASH:
		return true
	case TYPE_TAG_ARRAY:
		size := ti.ArrayFixedSize()
		if size != -1 {
			return TypeNeedsWrapper(ti.ParamType(0))
		}
		return true
	case TYPE_TAG_ERROR:
		panic("not implemented")
	case TYPE_TAG_INTERFACE:
		switch ti.Interface().Type() {
		case INFO_TYPE_CALLBACK, INFO_TYPE_ENUM, INFO_TYPE_FLAGS,
			INFO_TYPE_STRUCT, INFO_TYPE_UNION:
			return false
		}
		return true
	}
	return false
}
