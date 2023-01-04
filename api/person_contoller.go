package api

import (
	"encoding/json"
	"github.com/Unbel1evab7e/bank-interface/db/entity"
	"github.com/Unbel1evab7e/bank-interface/domain"
	"github.com/Unbel1evab7e/bank-interface/domain/dto"
	"github.com/Unbel1evab7e/bank-interface/integration/dadata"
	"github.com/Unbel1evab7e/bank-interface/service"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"io"
	"net/http"
)

type PersonController struct {
	DaDataClient  *dadata.DaDataClient
	PersonService *service.PersonService
}

func NewPersonController(personService *service.PersonService, daDataClient *dadata.DaDataClient, routerGroup *gin.RouterGroup) *PersonController {
	controller := &PersonController{
		DaDataClient:  daDataClient,
		PersonService: personService,
	}

	routerGroup.GET("/persons/suggestions", controller.GetSuggestions)
	routerGroup.GET("/persons", controller.Login)
	routerGroup.POST("/persons", controller.RegisterPerson)

	return controller
}

// GetSuggestions godoc
//
//	@Summary	Get All Suggestions of specified query
//	@Tags		Persons
//	@Produce	json
//	@Param		query	query		string	false	"Строка адреса"
//	@Success	200		{object}	domain.Response
//	@Failure	400		{object}	domain.Response
//	@Failure	500		{object}	domain.Response
//	@Router		/persons/suggestions [get]
func (pc *PersonController) GetSuggestions(c *gin.Context) {
	q := c.Request.URL.Query().Get("query")

	if q == "" {
		c.JSON(http.StatusBadRequest, domain.GenerateErrorResponse(`Parameter "query" is required`))
		return
	}

	suggestionResponse, err := pc.DaDataClient.GetSuggestions(q)

	if err != nil {
		log.Errorf("Failed to get suggestions", err)
		c.JSON(http.StatusInternalServerError, domain.GenerateErrorResponse(`Internal error`))
		return
	}

	c.JSON(http.StatusOK, domain.GenerateSuccessResponse[dadata.SuggestionResponse](suggestionResponse))
}

// RegisterPerson godoc
//
//	@Summary	Create new person
//	@Tags		Persons
//	@Produce	json
//	@Consume	json
//	@Param		person	body		dto.PersonDto	false	"Объект клиента"
//	@Success	200		{object}	domain.Response
//	@Failure	400		{object}	domain.Response
//	@Failure	500		{object}	domain.Response
//	@Router		/persons [post]
func (pc *PersonController) RegisterPerson(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.GenerateErrorResponse("Fail to obtain body"))
		return
	}

	var person dto.PersonDto

	err = json.Unmarshal(body, &person)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.GenerateErrorResponse("Fail to unmarshal body"))
		return
	}

	err = pc.PersonService.SavePerson(&person)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.GenerateErrorResponse("Fail to save person"))
		return
	}

	response := "OK"

	c.JSON(http.StatusOK, domain.GenerateSuccessResponse[string](&response))
}

// Login godoc
//
//	@Summary	Login and Auth person
//	@Tags		Persons
//	@Produce	json
//	@Param		phone		query		string	false	"Телефон клиента"
//	@Param		password	query		string	false	"Пароль клиента"
//	@Success	200			{object}	domain.Response
//	@Failure	400			{object}	domain.Response
//	@Failure	500			{object}	domain.Response
//	@Router		/persons [get]
func (pc *PersonController) Login(c *gin.Context) {
	ph := c.Request.URL.Query().Get("phone")

	if ph == "" {
		c.JSON(http.StatusBadRequest, domain.GenerateErrorResponse(`Parameter "phone" is required`))
		return
	}

	pass := c.Request.URL.Query().Get("password")

	if pass == "" {
		c.JSON(http.StatusBadRequest, domain.GenerateErrorResponse(`Parameter "password" is required`))
		return
	}

	person, err := pc.PersonService.GetPersonByPhoneAndPassword(ph, pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.GenerateErrorResponse("Fail to get person"))
		return
	}

	c.SetCookie("Authorize", "", 60000, "", "", true, false)
	c.JSON(http.StatusOK, domain.GenerateSuccessResponse[entity.Person](person))
}
