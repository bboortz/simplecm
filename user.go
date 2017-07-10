package main

import (
    osuser "os/user"
    "strconv"
	"syscall"
)


type name string
type id int
type dir string


const (
	ROOT_ID id = 0
	ROOT_NAME name = "root"
	ROOT_HOME dir = "/root"
)


type User interface {
	LogUser() 
	BecomeUser() User
	CheckRoot() 
	IsRoot() bool
}

type UserBuilder interface {
	SetUid(int) UserBuilder
	FromCurrentUser() UserBuilder
	FromUser(string) UserBuilder
	Build() User
}

type userBuilder struct {
	uid int
	gid int
	username string
	groupname string
	groupids []int
	homeDir string
}

func (b *userBuilder) SetUid(uid int) UserBuilder {
	b.uid = uid 
	return b
}

func (b *userBuilder) FromCurrentUser() UserBuilder {
	theUser, _ := osuser.Current()
	theGroup, _ := osuser.LookupGroupId( strconv.Itoa(b.gid) )
	b.uid, _ = strconv.Atoi(theUser.Uid)
	b.gid, _ = strconv.Atoi(theUser.Gid)
	b.username = theUser.Username
	b.groupname = theGroup.Name
	groupids, _ := theUser.GroupIds()
	b.groupids = convertStringArrayToIntArray( groupids )
	b.homeDir = theUser.HomeDir
	return b
}

func (b *userBuilder) FromUser(username string) UserBuilder {
	theUser, _ := osuser.Lookup(username)
	theGroup, _ := osuser.LookupGroupId( strconv.Itoa(b.gid) )
	b.uid, _ = strconv.Atoi(theUser.Uid)
	b.gid, _ = strconv.Atoi(theUser.Gid)
	b.username = theUser.Username
	b.groupname = theGroup.Name
	groupids, _ := theUser.GroupIds()
	b.groupids = convertStringArrayToIntArray( groupids )
	b.homeDir = theUser.HomeDir
	return b
}

func NewUser() UserBuilder {
	return &userBuilder{}
}


func (b *userBuilder) Build() User {
	return &user{
		uid: b.uid,
		gid: b.gid,
		username: b.username,
		groupname: b.groupname,
		groupids: b.groupids,
		homeDir: b.homeDir,
	}
}

type user struct {
	username string
	uid int
	groupname string
	gid int
	groupids []int
	homeDir string
}

func (u *user) LogUser() {
	log.Infof("user = %s, uid = %d, euid = %d, loginGroup = %s, gid = %d, egid = %d", u.username, u.uid, syscall.Geteuid(), u.groupname, u.gid, syscall.Getegid() )
}

func (u *user) BecomeUser() User {
	log.Debugf("becoming user <%s> with uid <%d> ...", u.username, u.uid )
	changeUser(u.uid, u.gid, u.groupids, u.homeDir)
	u.LogUser()

	return u
}

func (u *user) CheckRoot() {
        if ! u.IsRoot() {
                log.Error("You must run this program as user root!")
                programExit(1)
        }
}

func (u *user) IsRoot() bool {
        if u.uid != 0 {
		return false
        }
	return true
}


