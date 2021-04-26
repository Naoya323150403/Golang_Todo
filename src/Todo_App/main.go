package main

import (
	"Golang6/src/Todo_App/app/contollers"
	"Golang6/src/Todo_App/app/models"
	"fmt"
)

func main() {
	fmt.Println(models.Db)
	//user, _ := models.GetUserByEmail("naoya.kansai0403@gmail.com")
	//fmt.Println(user)
	contollers.StartMainServer()
	/*
		session, err := user.CreateSession()
		if err != nil{
			log.Fatalln(err)
		}
		fmt.Println(session)

		valid, _ := session.CheckSession()
		fmt.Println(valid)
	*/

}
