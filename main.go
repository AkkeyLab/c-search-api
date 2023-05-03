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

func getCompanyById(c *gin.Context) {
	id := c.Param("id")

	for _, cm := range companies {
		if cm.ID == id {
			c.IndentedJSON(http.StatusOK, cm)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Company not found"})
}

func postCompanies(c *gin.Context) {
	var createCompany company

	if err := c.BindJSON(&createCompany); err != nil {
		return
	}

	companies = append(companies, createCompany)
	c.IndentedJSON(http.StatusOK, createCompany)
}

func main() {
	router := gin.Default()
	router.GET("/companies", getCompanies)
	router.GET("/companies/:id", getCompanyById)
	router.POST("/companies", postCompanies)

	router.Run("localhost:8080")
}
