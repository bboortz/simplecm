package operatingsystem

import (
	"os"
	osuser "os/user"
	"strconv"
	"syscall"
)

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

func setUserPrivileges(username string, groupname string, path string) {
	user, err := osuser.Lookup(username)
	if err != nil {
		log.Fatalf("Failed to lookup username %s: %v", username, err)
	}
	/*
		group, err := osuser.LookupGroup(groupname)
		if err != nil {
			log.Fatalf("Failed to lookup groupname %s: %v", groupname, err)
		}
	*/

	uid, err := strconv.Atoi(user.Uid)
	// gid, err := strconv.Atoi(group.Gid)

	// set uid
	if syscall.Getuid() != uid {
		_, _, err := syscall.Syscall(syscall.SYS_SETUID, uintptr(uid), 0, 0)
		if err != 0 {
			log.Fatalf("Failed to syscall SYS_SETUID(%d): %v", uid, err)
		}
	}

	// set gid
	/*
		if syscall.Getgid() != gid {
			_, _, err := syscall.Syscall(syscall.SYS_SETGID, uintptr(gid), 0, 0)
			if err != 0 {
				log.Fatalf("Failed to syscall SYS_SETGID(%d): %v", gid, err)
			}
		}
	*/

	// chdir
	/*
		if path != "" {
			path = user.HomeDir
		}
		if err := os.Chdir(path); err != nil {
			log.Fatalf("Failed to Chdir to %q: %v", path, err)
		}
	*/
}
