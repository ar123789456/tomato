package main

import "tomato/server"

func main() {
	sqlDB := server.InitDB()
	defer sqlDB.Close()
}
