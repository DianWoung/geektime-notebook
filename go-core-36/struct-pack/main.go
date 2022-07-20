package main

import (
	"fmt"
	"struct-pack/helper"
	"time"
)

func main() {
	u := helper.UserInfo{}
	u.MakeUser("dialing", 33, time.Now())
	name := u.GetUserName()
	age := u.GetUserAge()
	fmt.Printf("user name is %s, age is %d \n", name, age)
}
