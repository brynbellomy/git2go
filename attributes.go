package git

/*
#include <git2.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type AttributeState uint32

const (
	AttributeUnspecified AttributeState = C.GIT_ATTR_UNSPECIFIED_T // The attribute has been left unspecified
	AttributeTrue        AttributeState = C.GIT_ATTR_TRUE_T        // The attribute has been set
	AttributeFalse       AttributeState = C.GIT_ATTR_FALSE_T       // The attribute has been unset
	AttributeValue       AttributeState = C.GIT_ATTR_VALUE_T       // This attribute has a value
)

type AttributeCheckFlag uint32

const (
	AttributeCheckFileThenIndex AttributeCheckFlag = C.GIT_ATTR_CHECK_FILE_THEN_INDEX
	AttributeCheckIndexThenFile AttributeCheckFlag = C.GIT_ATTR_CHECK_INDEX_THEN_FILE
	AttributeCheckIndexOnly     AttributeCheckFlag = C.GIT_ATTR_CHECK_INDEX_ONLY

	// Check attribute flags: Using the system attributes file.
	//
	// Normally, attribute checks include looking in the /etc (or system
	// equivalent) directory for a `gitattributes` file.  Passing this
	// flag will cause attribute checks to ignore that file.
	AttributeCheckNoSystem AttributeCheckFlag = C.GIT_ATTR_CHECK_NO_SYSTEM
)

type Attribute string

// GIT_EXTERN(int) git_attr_get(
//     const char **value_out,
//     git_repository *repo,
//     uint32_t flags,
//     const char *path,
//     const char *name);
func (r *Repository) GetAttribute(flags AttributeCheckFlag, path string, name string) (Attribute, error) {
	var out *C.char

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	ret := C.git_attr_get(&out, r.ptr, C.uint32_t(flags), cPath, cName)
	runtime.KeepAlive(r)
	if ret != 0 {
		return "", MakeGitError(ret)
	}

	return Attribute(C.GoString(out)), nil
}

func (r *Repository) AddAttributeMacro(name string, values string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cValues := C.CString(values)
	defer C.free(unsafe.Pointer(cValues))

	ret := C.git_attr_add_macro(r.ptr, cName, cValues)
	runtime.KeepAlive(r)
	if ret != 0 {
		return MakeGitError(ret)
	}
	return nil
}

// void git_attr_cache_flush(git_repository *repo);
func (r *Repository) FlushAttributeCache() {
	C.git_attr_cache_flush(r.ptr)
	runtime.KeepAlive(r)
}

func (attr Attribute) State() AttributeState {
	cAttr := C.CString(string(attr))
	defer C.free(unsafe.Pointer(cAttr))

	return AttributeState(C.git_attr_value(cAttr))
}

func (attr Attribute) HasValue() bool {
	return attr.State() == AttributeValue //C.git_attr_value(cAttr) == C.GIT_ATTR_VALUE_T
}

func (attr Attribute) IsFalse() bool {
	return attr.State() == AttributeFalse
}

func (attr Attribute) IsTrue() bool {
	return attr.State() == AttributeTrue
}

func (attr Attribute) IsUnspecified() bool {
	return attr.State() == AttributeUnspecified
}
