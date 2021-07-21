package main

import (
	"repos.tdctechsupport.com/t2015014/user_management_server/route"
)

func main() {

	router := route.Init()
	router.Start(":8082")
}
