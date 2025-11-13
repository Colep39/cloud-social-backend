package main

import (
  "log"
  "net/http"
  "os"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.POST("/register", registerHandler)
  r.POST("/login", loginHandler)
  r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

  addr := ":8081"
  if p := os.Getenv("PORT"); p != "" { addr = ":" + p }
  log.Printf("auth service listening on %s", addr)
  _ = r.Run(addr)
}

func registerHandler(c *gin.Context) {
  // TODO: parse payload, hash password, insert into users table
  c.JSON(201, gin.H{"ok": true})
}

func loginHandler(c *gin.Context) {
  // TODO: verify password, issue JWT
  c.JSON(200, gin.H{"token": "fake.jwt.token"})
}
