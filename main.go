package main

import (
	"golang-challenge/database"
	server2 "golang-challenge/server"
)

func main() {
	database.StartDB()

	server := server2.NewServer()

	server.Run()
}
