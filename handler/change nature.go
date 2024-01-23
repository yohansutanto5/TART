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

type ChangeNatureHandler struct {
	CRUDHandler
	ChangeNatureService service.ChangeNatureService
}

func NewChangeNatureHandler(db *db.DataStore) ChangeNatureHandler {
	svc := service.NewChangeNatureService(db)
	h := ChangeNatureHandler{
		ChangeNatureService: svc,
	}
	return h
}

func (h *ChangeNatureHandler) GetList(c *gin.Context) {
	// Call Service
	result, err := h.ChangeNatureService.GetList()

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (h *ChangeNatureHandler) Insert(c *gin.Context) {
	// Cast data from request
	var data model.AddChangeNatureIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Construct ChangeNature Model with the request data
	newChangeNature := &model.ChangeNature{}
	newChangeNature.PopulateFromDTOInput(data)
	// Call create service
	err := h.ChangeNatureService.Insert(newChangeNature)

	// Construct Response
	if err != nil {
		err.GenerateReponse(util.GetTransactionID(c))
		c.JSON(err.Status, err.Response)
	} else {
		c.JSON(http.StatusCreated, constanta.SuccessMessage)
	}
}
