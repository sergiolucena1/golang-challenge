package main

import server2 "golang-challenge/server"

func main() {
	server := server2.NewServer()

	server.Run()
}
