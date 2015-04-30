package getgroup

/*
#include <stdlib.h>
#include <errno.h>
#include <sys/types.h>
#include <grp.h>
#include <unistd.h>

struct group * nametogrp(char *name, char **buf) {
    size_t bufsize = sysconf(_SC_GETGR_R_SIZE_MAX);
    if (bufsize == -1) {
        bufsize = 16384;
    }

    *buf = malloc(bufsize);
    if (*buf == NULL) {
        errno = ENOMEM;
        return NULL;
    }

    struct group *grp;
    struct group *result;

    grp = malloc(sizeof(struct group));
    if (grp == NULL) {
        errno = ENOMEM;
        return NULL;
    }

    int ret = getgrnam_r(name, grp, *buf, bufsize, &result);
    free(name);
    if (result == NULL) {
        if (ret != 0) {
            errno = ret;
        }
        return NULL;
    }
    return result;
}

static inline int getMemLength(char **mem) {
	int len;
	if (mem == NULL)
		return 0;
	for (len = 0; *mem; len++, mem++);
	return len;
}
*/
import "C"
import "unsafe"
import "fmt"
import "reflect"

//NewGroupFromName returns a new Group given a username
//or an error if any occurred
func NewGroupFromName(name string) (*Group, error) {
	var buf = new(*C.char)
	grp, err := C.nametogrp(C.CString(name), buf)
	defer C.free(unsafe.Pointer(*buf))
	defer C.free(unsafe.Pointer(grp))

	if err != nil {
		return nil, err
	}
	if grp == nil {
		return nil, fmt.Errorf("gid does not exist")
	}
	length := int(C.getMemLength(grp.gr_mem))
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(grp.gr_mem)),
		Len: length,
		Cap: length,
	}
	return &Group{
		Name: C.GoString(grp.gr_name),
		Passwd: C.GoString(grp.gr_passwd),
		Gid: uint32(grp.gr_gid),
		Mem: *(*[]string)(unsafe.Pointer(&hdr)),
	}, nil
}
