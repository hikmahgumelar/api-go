package main

import "github.com/gin-gonic/gin" 
import "fmt"
var a string = "21"

func main(){
     r := gin.Default()

     // Routing GET 
      r.GET("/", func(c *gin.Context){
          c.JSON(200, gin.H{
               "Pesan":"Welcome",
            })

      })
      r.GET("/room", func(c *gin.Context){
          c.JSON(200, gin.H{
               
               "Pesan":fmt.Sprintf("Welcome to room %s",a),
						 })

      })
			//Run server
      r.Run()
  }
