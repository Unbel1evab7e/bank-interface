package api

import (
	"bank-interface/domain"
	"bank-interface/integration/dadata"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"
	"net/http"
)

type PersonController struct {
	DaDataClient *dadata.DaDataClient
}

func NewPersonController(daDataClient *dadata.DaDataClient, routerGroup *gin.RouterGroup) *PersonController {
	controller := &PersonController{DaDataClient: daDataClient}

	routerGroup.GET("/persons/suggestions", controller.GetSuggestions)

	return controller
}

// GetSuggestions godoc
// @Summary      Get All Suggestions of specified query
// @Tags         Persons
// @Produce      json
// @Param 		 query query	string	false "Строка адреса"
// @Success      200  {object}  domain.Response
// @Failure		 400  {object}  domain.Response
// @Failure      500  {object}  domain.Response
// @Router /persons/suggestions [get]
func (pc PersonController) GetSuggestions(c *gin.Context) {
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
