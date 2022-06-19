package main

type user struct {
	ID     int64
	Name   string
	Avatar string
}

// func getUserInfo() *user {
// 	return &user{ID: 13746554, Name: "xubeiping", Avatar: "https://hellolinux.xyz"}
// }

// func getUserInfo2(u *user) *user {
// return u
// }

func getUserInfo3(u user) *user {
	return &u
}

func main() {
	// _ = getUserInfo()
	_ = getUserInfo3(user{ID: 13746554, Name: "xubeiping", Avatar: "https://hellolinux.xyz"})
}
