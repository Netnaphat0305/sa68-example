package main

import (
	"fmt"
	"example.com/sa68-example/config"
)

func main() {
	// เปิดการเชื่อมต่อฐานข้อมูล
	config.ConnectionDB()

	// Setup database และ seed ข้อมูล
	config.SetupDatabase()

	fmt.Println("Database initialized and seeded successfully.")
}
