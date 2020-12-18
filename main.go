package main

import (
	"fmt"
	"zapood-webapi/restapi"
	"zapood/config"
)

func main() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		r := restapi.RunApi(db)
		//running
		r.Run(":8080")
	}
}
