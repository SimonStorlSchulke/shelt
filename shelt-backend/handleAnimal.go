package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

// eg: http://localhost:8080/animal/1?fields=name,gender,shoulderheight
func GetAnimalFieldsById(c *gin.Context) {

	id := c.Param("id")

	strapiQuery, descriptionLength := BuildFieldsQuery(c)

	url := "animals/" + id + strapiQuery

	body, err := StrapiGet(url)

	logErr(err)

	var strapiData map[string]interface{}

	json.Unmarshal(body, &strapiData)
	data := strapiData["data"].(map[string]interface{})

	fmt.Println("DATA:", data)

	description := data["attributes"].(map[string]interface{})["Description"].(string)
	mnLen := int(math.Min(float64(descriptionLength), float64(len(description))))
	data["attributes"].(map[string]interface{})["Description"] = description[0:mnLen]

	c.JSON(http.StatusOK, data["attributes"])
}

// eg: http://localhost:8080/animals?fields=name,gender
func GetAnimalCollection(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:4200")

	strapiQuery, descriptionLength := BuildFieldsQuery(c)

	url := "animals" + strapiQuery

	body, err := StrapiGet(url)

	logErr(err)

	var strapiData map[string]interface{}

	json.Unmarshal(body, &strapiData)
	fmt.Println(string(body))
	data := strapiData["data"].([]interface{})

	var returnData []interface{} = make([]interface{}, len(data))

	for i, d := range data {
		id := d.(map[string]interface{})["id"]
		returnData[i] = d.(map[string]interface{})["attributes"]
		fields := returnData[i].(map[string]interface{})

		if fields["Description"] != nil && descriptionLength != 0 {
			description := fields["Description"].(string)
			mnLen := int(math.Min(float64(descriptionLength), float64(len(description))))
			fields["Description"] = description[0:mnLen]
		}

		fields["id"] = id
	}

	c.JSON(http.StatusOK, returnData)
}
