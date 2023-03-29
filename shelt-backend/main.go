package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func logErr(err error) {
	if err != nil {
		log.Println("ERROR:", err)
	}
}

type M map[string]interface{}

var strapiKey string = "37fe1534ac2d292a85373fa183b18062a73c32cc575a01bee7e7a8a4599bbc3113df6e31df96a58fb4884b76508ea6a3da7735bfced2b5600ee128ebe29e2c628f503ba9c3800b14dd91cef8d119413e84573e4b36323dcdb35b00f83acbdd688cc8dc664c6c9a1de86086d6cb58d98bbe83770db708dbfdafd64d947e2b3974"

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
