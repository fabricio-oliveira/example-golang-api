package main

import (
	"local/descomplica-company/api/conf"
)

func main() {

	db, error := conf.InitDB()
	if error != nil {
		panic(error)
	}

	conf.InitHandle(db)
}
