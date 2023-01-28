package services

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"example.com/libr34/models"
	"github.com/gin-gonic/gin"
)

const BaseUrl string = "https://api.rule34.xxx/index.php?"

func AttachServices(r *gin.Engine) *gin.Engine {
	r.GET("/", search)
	r.GET("/post", post)
	r.GET("/posts", posts)
	r.GET("/autocomplete", autocomplete)

	return r
}

func search(c *gin.Context) {
	c.JSON(200, "test")
}

func posts(c *gin.Context) {
	var posts models.Posts

	tags := c.Query("tags")

	response, err := http.Get(BaseUrl + "page=dapi&s=post&q=index&tags=" + tags + "&pid=1")

	if err != nil {
		fmt.Println("Error making GET request: ", err)
		c.JSON(500, "test")
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		c.JSON(401, "test")
		return
	}

	xml.Unmarshal(body, &posts)

	jsonData, err := json.Marshal(posts)

	c.JSON(200, string(jsonData))
}

func post(c *gin.Context) {
	var posts models.Posts

	id := c.Query("id")

	response, err := http.Get(BaseUrl + "page=dapi&s=post&q=index&id=" + id)

	if err != nil {
		fmt.Println("Error making GET request: ", err)
		c.JSON(500, "test")
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		c.JSON(401, "test")
		return
	}

	xml.Unmarshal(body, &posts)

	jsonData, err := json.Marshal(posts)

	c.JSON(200, string(jsonData))
}
