package main

import (
	"app/handler"

	"github.com/gin-gonic/gin"
)

// Rest API
func setupRoutes() *gin.Engine {
	r := gin.New()
	// Setup Middleware
	r.Use(middleware, gin.Recovery())

	// Initiate all Handler and dependency
	userHandler := handler.NewUserHandler(database)
	userProfilehandler := handler.NewUserProfileHandler(database)
	systemHander := handler.NewSystemHandler(database)
	ticketHandler := handler.NewTicketHandler(database)
	changeNaturehandler := handler.NewChangeNatureHandler(database)

	// Define The route Path
	// ---- System API ---
	r.GET("/health", systemHander.GetSystemHealth)

	// ---- User API ---
	r.POST("/user", userHandler.Insert)
	r.GET("/user", userHandler.GetList)

	// ---- UserProfile API ---
	r.POST("/userprofile", userProfilehandler.Insert)
	r.GET("/userprofile", userProfilehandler.GetList)

	// ---- Ticket API ---
	r.POST("/ticket", ticketHandler.Insert)
	r.GET("/ticket", ticketHandler.GetList)

	// ---- UserProfile API ---
	r.POST("/changenature", changeNaturehandler.Insert)
	r.GET("/changenature", changeNaturehandler.GetList)

	return r
}
