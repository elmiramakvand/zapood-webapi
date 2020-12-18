package main

import (
	"fmt"
	"zapood-webapi/restapi"
	"zapood/config"
)

func main() {
	db, err := config.GetDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		r := restapi.RunApi(db)
		//running
		r.Run()

	}
}
