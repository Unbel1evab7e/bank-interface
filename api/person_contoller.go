package api

import (
	"encoding/json"
	"github.com/Unbel1evab7e/bank-interface/domain"
	"github.com/Unbel1evab7e/bank-interface/domain/dto"
	"github.com/Unbel1evab7e/bank-interface/integration/dadata"
	"github.com/Unbel1evab7e/bank-interface/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"io"
	"net/http"
)

type PersonController struct {
	DaDataClient  *dadata.DaDataClient
	PersonService *service.PersonService
	JwtMiddleware *jwt.GinJWTMiddleware
}

func NewPersonController(personService *service.PersonService,
	daDataClient *dadata.DaDataClient,
	routerGroup *gin.RouterGroup,
	jwtMiddleware *jwt.GinJWTMiddleware) *PersonController {
	controller := &PersonController{
		DaDataClient:  daDataClient,
		PersonService: personService,
		JwtMiddleware: jwtMiddleware,
	}

	auth := routerGroup.Group("/auth")

	auth.Use(jwtMiddleware.MiddlewareFunc())

	{
		auth.GET("/addresses", controller.GetSuggestions)
	}

	routerGroup.POST("/persons/login", controller.Login)
	routerGroup.POST("/persons/logout", controller.Logout)
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
//	@Router		/auth/addresses [get]
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
//	@Param		login	body		dto.LoginDto	false	"Объект логина"
//	@Success	200		{object}	domain.Response
//	@Failure	400		{object}	domain.Response
//	@Failure	500		{object}	domain.Response
//	@Router		/persons/login [post]
func (pc *PersonController) Login(c *gin.Context) {
	pc.JwtMiddleware.LoginHandler(c)
}

// Logout godoc
//
//	@Summary	Logout person
//	@Tags		Persons
//	@Success	200	{object}	domain.Response
//	@Failure	400	{object}	domain.Response
//	@Failure	500	{object}	domain.Response
//	@Router		/persons/logout [post]
func (pc *PersonController) Logout(c *gin.Context) {
	pc.JwtMiddleware.LogoutHandler(c)
}
