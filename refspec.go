package git

/*
#include <git2.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type Refspec struct {
	remote *Remote
	ptr    *C.git_refspec
}

func (rs *Refspec) Transform(refname string) (string, error) {
	crefname := C.CString(refname)
	defer C.free(unsafe.Pointer(crefname))

	out := C.git_buf{}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := C.git_refspec_transform(&out, rs.ptr, crefname)
	runtime.KeepAlive(rs)
	if ret < 0 {
		return "", MakeGitError(ret)
	}
	defer C.git_buf_free(&out)

	return C.GoString(out.ptr), nil
}

func (rs *Refspec) RTransform(refname string) (string, error) {
	crefname := C.CString(refname)
	defer C.free(unsafe.Pointer(crefname))

	out := C.git_buf{}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ret := C.git_refspec_rtransform(&out, rs.ptr, crefname)
	runtime.KeepAlive(rs)
	if ret < 0 {
		return "", MakeGitError(ret)
	}
	defer C.git_buf_free(&out)

	return C.GoString(out.ptr), nil
}

func (rs *Refspec) String() string {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	cstr := C.git_refspec_string(rs.ptr)
	runtime.KeepAlive(rs)
	defer C.free(unsafe.Pointer(cstr))

	return C.GoString(cstr)
}

func (rs *Refspec) Direction() ConnectDirection {
	dir := C.git_refspec_direction(rs.ptr)
	runtime.KeepAlive(rs)
	return ConnectDirection(dir)
}

func (rs *Refspec) SourceMatches(refname string) bool {
	crefname := C.CString(refname)
	defer C.free(unsafe.Pointer(crefname))

	ret := C.git_refspec_src_matches(rs.ptr, crefname)
	runtime.KeepAlive(rs)
	if ret == 1 {
		return true
	}
	return false
}

func (rs *Refspec) DestinationMatches(refname string) bool {
	crefname := C.CString(refname)
	defer C.free(unsafe.Pointer(crefname))

	ret := C.git_refspec_dst_matches(rs.ptr, crefname)
	runtime.KeepAlive(rs)
	if ret == 1 {
		return true
	}
	return false
}
