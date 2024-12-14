package main

import (
	"C"
	"fmt"
	"encoding/json"
)

type User struct {
	User string
	Email string
}

//export login
func login(str *C.char) *C.char {
	user  := C.GoString(str)
	email := fmt.Sprintf("%s@example.com", user)
	obj   := &User{User: user, Email: email}

	json, err := json.Marshal(obj)

	if (err == nil) {
		return C.CString(string(json))
	} else {
		fmt.Println(err)
		return nil
	}
}

func main() {}
