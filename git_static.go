// +build static

package git

/*
#cgo CFLAGS: -I${SRCDIR}/vendor/libgit2/include
#cgo linux LDFLAGS: -L${SRCDIR}/libgit2-static/linux -lgit2
#cgo linux pkg-config: --static ${SRCDIR}/libgit2-static/linux/libgit2.pc
#cgo darwin LDFLAGS: -L${SRCDIR}/libgit2-static/darwin -lgit2
#cgo darwin pkg-config: --static ${SRCDIR}/libgit2-static/darwin/libgit2.pc
#cgo windows LDFLAGS: -L${SRCDIR}/libgit2-static/windows -lgit2 -lwinhttp
#include <git2.h>

#if LIBGIT2_VER_MAJOR != 0 || LIBGIT2_VER_MINOR != 27
# error "Invalid libgit2 version; this git2go supports libgit2 v0.27"
#endif

*/
import "C"
