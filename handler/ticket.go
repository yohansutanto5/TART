package handler

import (
	"app/constanta"
	"app/db"
	"app/model"
	"app/pkg/util"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	CRUDHandler
	TicketService service.TicketService
}

func NewTicketHandler(db *db.DataStore) TicketHandler {
	svc := service.NewTicketService(db)
	h := TicketHandler{
		TicketService: svc,
	}
	return h
}

func (h *TicketHandler) GetList(c *gin.Context) {
	// Call Service
	result, err := h.TicketService.GetList()

	// Construct Response
	if err != nil {
		c.Errors = append(c.Errors, err.GenerateReponse(util.GetTransactionID(c)))
		c.JSON(err.Status, err.Response)
	} else {
		// Construct DTO out
		var response []model.GetTicketOut
		for _, Ticket := range result {
			response = append(response, Ticket.ConstructGetTicketOut())
		}
		c.JSON(http.StatusOK, response)
	}
}

func (h *TicketHandler) Insert(c *gin.Context) {
	// Cast data from request
	var data model.AddTicketIn
	if errx := c.ShouldBindJSON(&data); errx != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errx.Error()})
		return
	}
	// Construct Ticket Model with the request data
	newTicket := &model.Ticket{}
	newTicket.PopulateFromDTOInput(data)
	// Call create service
	err := h.TicketService.Insert(newTicket)

	// Construct Response
	if err != nil {
		c.Errors = append(c.Errors, err.GenerateReponse(util.GetTransactionID(c)))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}
