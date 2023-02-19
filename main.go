package main

  import (
    "webScraperBackend/controllers"
    "webScraperBackend/initializers"
    "github.com/gin-gonic/gin"
  )

  func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDatabase()
  }

  func main() {
    r := gin.Default()
    r.GET("/", controllers.SitesCreate)
    r.POST("/signup", controllers.SignUp)
    r.POST("/login", controllers.Login)
    r.Run()
  }