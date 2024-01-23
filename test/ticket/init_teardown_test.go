package user_test

import (
	"app/cmd/config"
	"app/db"
	"app/handler"
	"app/model"
	"app/service"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// Create testing Contenxt
var ctx *gin.Context
var configuration config.Configuration
var database *db.DataStore
var ticketService service.TicketService
var ticketHandler handler.TicketHandler

func TestMain(m *testing.M) {
	mode := os.Getenv("GIN_MODE")
	configfolder := os.Getenv("config")
	configuration = config.Load(mode, configfolder)
	database = db.NewDatabase(configuration)
	ticketService = service.NewTicketService(database)
	ticketHandler = handler.NewTicketHandler(database)

	// Create a mock HTTP request for testing
	req, _ := http.NewRequest("GET", "/mockurl", nil)
	w := httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Request = req

	// Input Init data
	profile := model.UserProfile{
		Name: "admin",
	}
	database.Db.Create(&profile)
	// Run tests
	exitCode := m.Run()

	// Cleanup resources, close the database connection, etc.
	if err := database.Db.Exec("TRUNCATE TABLE regressions;").Error; err != nil {
		panic(fmt.Sprintf("Failed to truncate table: %v", err))
	}

	if err := database.Db.Exec("ALTER TABLE regressions AUTO_INCREMENT = 1;").Error; err != nil {
		panic(fmt.Sprintf("Failed to reset auto-increment sequence: %v", err))
	}
	if err := database.Db.Exec("TRUNCATE TABLE performances;").Error; err != nil {
		panic(fmt.Sprintf("Failed to truncate table: %v", err))
	}

	if err := database.Db.Exec("ALTER TABLE performances AUTO_INCREMENT = 1;").Error; err != nil {
		panic(fmt.Sprintf("Failed to reset auto-increment sequence: %v", err))
	}
	if err := database.Db.Exec("TRUNCATE TABLE uats;").Error; err != nil {
		panic(fmt.Sprintf("Failed to truncate table: %v", err))
	}

	if err := database.Db.Exec("ALTER TABLE uats AUTO_INCREMENT = 1;").Error; err != nil {
		panic(fmt.Sprintf("Failed to reset auto-increment sequence: %v", err))
	}
	if err := database.Db.Exec("DELETE FROM tickets;").Error; err != nil {
		panic(fmt.Sprintf("Failed to truncate table: %v", err))
	}

	if err := database.Db.Exec("ALTER TABLE tickets AUTO_INCREMENT = 1;").Error; err != nil {
		panic(fmt.Sprintf("Failed to reset auto-increment sequence: %v", err))
	}

	if err := database.Db.Exec("TRUNCATE TABLE chaos;").Error; err != nil {
		panic(fmt.Sprintf("Failed to truncate table: %v", err))
	}

	if err := database.Db.Exec("ALTER TABLE chaos AUTO_INCREMENT = 1;").Error; err != nil {
		panic(fmt.Sprintf("Failed to reset auto-increment sequence: %v", err))
	}

	if err := database.Db.Exec("TRUNCATE TABLE sonars;").Error; err != nil {
		panic(fmt.Sprintf("Failed to truncate table: %v", err))
	}

	if err := database.Db.Exec("ALTER TABLE sonars AUTO_INCREMENT = 1;").Error; err != nil {
		panic(fmt.Sprintf("Failed to reset auto-increment sequence: %v", err))
	}

	if err := database.Db.Exec("TRUNCATE TABLE unit_tests;").Error; err != nil {
		panic(fmt.Sprintf("Failed to truncate table: %v", err))
	}

	if err := database.Db.Exec("ALTER TABLE unit_tests AUTO_INCREMENT = 1;").Error; err != nil {
		panic(fmt.Sprintf("Failed to reset auto-increment sequence: %v", err))
	}

	os.Exit(exitCode)

}
