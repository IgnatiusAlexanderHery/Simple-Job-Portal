package main

import (
	"log"

	"Simple-Job-Portal/auth"
	"Simple-Job-Portal/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	csrf "github.com/srbry/gin-csrf"
)

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	csrfManager := &csrf.DefaultCSRFManager{
		Secret: "secret",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}
	router.Use(sessions.Sessions("X-CSRF-Token", store))
	router.Use(func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Next()
	})

	router.POST("/login", func(c *gin.Context) {
		controller.Login(c, csrfManager)
	})
	router.POST("/register", controller.Register)
	router.POST("/logout", controller.Logout)
	router.GET("/profile", controller.GetUserData)

	talentRouter := router.Group("/talent")
	talentRouter.Use(auth.AuthMiddlewareTalent()).Use(csrfManager.Middleware())
	{
		talentRouter.GET("/jobs", controller.ViewJobs)
		talentRouter.POST("/jobs/:jobID/apply", controller.ApplyForJob)
		talentRouter.GET("/jobs/:jobID", controller.ViewJobDetail)
		talentRouter.GET("/applications", controller.ViewApplications)
		talentRouter.GET("/jobs/:jobID/applications", controller.ViewApplications)
	}

	employerRouter := router.Group("/employer")
	employerRouter.Use(auth.AuthMiddlewareEmployer()).Use(csrfManager.Middleware())
	{
		employerRouter.GET("/jobs", controller.ViewJobsByEmployer)
		employerRouter.POST("/jobs", controller.CreateJob)
		employerRouter.PUT("/jobs/:jobID", controller.UpdateJob)
		employerRouter.GET("/jobs/:jobID", controller.ViewJobDetail)
		employerRouter.GET("/jobs/:jobID/applications", controller.ViewApplications)
		employerRouter.GET("/applications", controller.ViewApplications)
		employerRouter.POST("/applications/:applicationID", controller.ProcessApplication)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
