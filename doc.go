/*
getgroup is a thin wrapper around getgrnam_r and getgrgid_r.
I wrote this package because the syscall package doesn't have any way to map a id to a groupname.
At this point the API shouldn't change, but I won't make any promises yet.
*/
package getgroup
