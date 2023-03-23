package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const STRAPI_BASE_URL = "http://localhost:1337/api/"

// StrapiGet
func StrapiGet(url string) ([]byte, error) {
	fmt.Println("Request api:", url)
	req, err := http.NewRequest("GET", STRAPI_BASE_URL+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+strapiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

func BuildFieldsQuery(c *gin.Context) string {

	fieldsQuery := c.Query("fields")
	fieldsStrapiQuery := ""

	if fieldsQuery != "" {
		fields := strings.Split(fieldsQuery, ",")

		fieldsStrapiQuery += "?"

		for i, field := range fields {
			fieldsStrapiQuery += fmt.Sprintf("fields[%d]=%s", i, field)
			if i != len(fields)-1 {
				fieldsStrapiQuery += "&"
			}
		}
	}

	return fieldsStrapiQuery
}
