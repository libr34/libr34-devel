package services

import (
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
	r.GET("/posts", posts)

	return r
}

func search(c *gin.Context) {
	c.JSON(200, "test")
}

func posts(c *gin.Context) {
	var posts models.Posts

	tags := c.Query("tags")

	response, err := http.Get(BaseUrl + "page=dapi&s=post&q=index&tags=" + tags + "&limit=5")

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

	fmt.Println(string(body))

	xml.Unmarshal(body, &posts)

	fmt.Println(posts.PostS[1].Tags)

	c.JSON(200, "test")
}

func post(c *gin.Context) {

}
