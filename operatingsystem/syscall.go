package operatingsystem

import (
	"os"
	"syscall"
)

// ChangeUser let the current user switch to another user
func ChangeUser(uid int, gid int, groupids []int, path string) {
	var err error

	// set gid
	if syscall.Getgid() != gid {
		err = syscall.Setregid(gid, 0)
		if err != nil {
			log.Fatalf("Failed to syscall SYS_SETREGID(%d): %v", gid, err)
		}
		err = syscall.Setgroups(groupids)
		if err != nil {
			log.Fatalf("Failed to syscall SYS_SETGROUPS(%d): %v", gid, err)
		}
	}

	// set uid
	if syscall.Getuid() != uid && syscall.Geteuid() != uid {
		err = syscall.Setreuid(uid, 0)
		if err != nil {
			log.Fatalf("Failed to syscall SYS_SETREUID(%d): %v", uid, err)
		}
	}

	// chdir
	if err := os.Chdir(path); err != nil {
		log.Fatalf("Failed to Chdir to %q: %v", path, err)
	}
}

// LinkFile links the file newpath to the oldpath
func LinkFile(oldpath string, newpath string) error {
	if !fileExists(newpath) {
		return syscall.Link(oldpath, newpath)
	}

	return nil
}
