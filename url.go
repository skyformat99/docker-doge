package main

import (
	"docker-doge/handler"
	"docker-doge/middleware"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func URL(r *gin.Engine) {

	e := middleware.GetAuthzInstance()
	middleware.NewJwtAuthorizer(e)                 // jwt权限校验器
	jwtMiddleWare := middleware.NewJwtMiddleWare() // jwt中间件
	r.Use(gin.Logger())                            // 日志处理
	r.Use(gin.Recovery())                          // 500不处理
	// API
	r.POST("/login", jwtMiddleWare.LoginHandler)
	r.POST("/register", handler.RegisterHandler)
	auth := r.Group("/auth")
	auth.Use(jwtMiddleWare.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			c.JSON(200, gin.H{
				"username": claims["id"],
				"text":     "Hello World.",
			})
		})
		auth.GET("/refresh_token", jwtMiddleWare.RefreshHandler)
	}
}