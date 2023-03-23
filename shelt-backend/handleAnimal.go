package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// eg: http://localhost:8080/animal/1?fields=name,gender,shoulderheight
func GetAnimalFieldsById(c *gin.Context) {

	id := c.Param("id")

	url := "animals/" + id + BuildFieldsQuery(c)

	body, err := StrapiGet(url)

	logErr(err)

	var strapiData map[string]interface{}

	json.Unmarshal(body, &strapiData)
	data := strapiData["data"].(map[string]interface{})

	c.JSON(http.StatusOK, data)
}

// eg: http://localhost:8080/animals?fields=name,gender
func GetAnimalCollection(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")

	url := "animals" + BuildFieldsQuery(c)

	body, err := StrapiGet(url)

	logErr(err)

	var strapiData map[string]interface{}

	json.Unmarshal(body, &strapiData)
	data := strapiData["data"].([]interface{})

	c.JSON(http.StatusOK, data)
}
