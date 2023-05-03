package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var companies = []company{
	{ID: "1", Name: "AkkeyLab, inc."},
	{ID: "2", Name: "Apple, inc."},
}

func getCompanies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, companies)
}

func main() {
	router := gin.Default()
	router.GET("/companies", getCompanies)

	router.Run("localhost:8080")
}
