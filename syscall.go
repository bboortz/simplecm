package main

import (
	"syscall"
//	"os"
	osuser "os/user"
	"strconv"
)

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
