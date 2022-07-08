package main

import(
	// "Go-Commerce/controllers"
	// "Go-Commerce/database"
	// "Go-Commerce/middleware"
	// "Go-Commerce/routes"
	//"log"
	"os"
	//"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8000"
	}
	
}
