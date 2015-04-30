package getgroup

//Group contains grop information in a format similar to grp.h's struct group
type Group struct {
	Name string
	Passwd string
	Gid uint32
	Mem []string
}

