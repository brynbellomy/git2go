// +build static

package git

/*
#cgo CFLAGS: -I${SRCDIR}/vendor/libgit2/include
#cgo linux LDFLAGS: -L${SRCDIR}/libgit2-static/linux-latest -lgit2
#cgo linux pkg-config: --static ${SRCDIR}/libgit2-static/linux-latest/libgit2.pc
#cgo darwin LDFLAGS: -L${SRCDIR}/libgit2-static/macos-latest -lgit2
#cgo darwin pkg-config: --static ${SRCDIR}/libgit2-static/macos-latest/libgit2.pc
#cgo windows LDFLAGS: -L${SRCDIR}/libgit2-static/darwin-latest -lgit2 -lwinhttp
#include <git2.h>

#if LIBGIT2_VER_MAJOR != 0 || LIBGIT2_VER_MINOR != 27
# error "Invalid libgit2 version; this git2go supports libgit2 v0.27"
#endif

*/
import "C"
