package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func logErr(err error) {
	if err != nil {
		log.Println("ERROR:", err)
	}
}

type M map[string]interface{}

var strapiKey string = os.Args[1]

func main() {
	r := gin.Default()
	r.GET("/animal/:id", GetAnimalFieldsById)
	r.GET("/animals", GetAnimalCollection)
	r.GET("/article/:id", GetArticle)
	r.GET("/animal-view", GetAnimalView)
	r.GET("/animal-collection-view", GetAnimalListView)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func mainN() {
	data := make(map[string]interface{})
	data["Key"] = "TEST"
	str, err := RenderTemplate("<h1>{{.Key}}</h1>", data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	//r := gin.Default()
	//r.GET("/animal/:id", GetAnimalFieldsById)
	//r.GET("/animals", GetAnimalCollection)
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
