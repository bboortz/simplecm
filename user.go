package main

import (
    osuser "os/user"
	"os"
    "strconv"
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
	BecomeUser() 
	CheckRoot() 
	IsRoot() bool
}

type UserBuilder interface {
	SetUid(int) UserBuilder
	FromCurrentUser() UserBuilder
	Build() User
}

type userBuilder struct {
	username string
	uid int
	groupname string
	gid int
	homeDir string
}

func (b *userBuilder) SetUid(uid int) UserBuilder {
	b.uid = uid 
	return b
}

func (b *userBuilder) FromCurrentUser() UserBuilder {
	currentUser, _ := osuser.Current()
	//uid := strconv.Atoi(currentUser.Uid)
	// b.uid = strconv.Atoi( currentUser.Uid )
	b.uid = os.Getuid()
	b.gid = os.Getgid()
	log.Debugf("uid: <%d>", b.uid)
	log.Debugf("gid: <%d>", b.gid)
	b.username = currentUser.Username
	currentGroup, _ := osuser.LookupGroupId( strconv.Itoa(b.gid) )
	b.groupname = currentGroup.Name
	b.homeDir = currentUser.HomeDir
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
		homeDir: b.homeDir,
	}
}

type user struct {
	username string
	uid int
	groupname string
	gid int
	homeDir string
}

func (u *user) LogUser() {
	log.Infof("user = %s, uid = %d, loginGroup = %s, gid = %d", u.username, u.uid, u.groupname, u.gid )
}

func (u *user) BecomeUser() {
	log.Debugf("becoming user <%s> with uid <%d> ...", u.username, u.uid )
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


