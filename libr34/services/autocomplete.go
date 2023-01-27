package services

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func autocomplete(c *gin.Context) {
	fmt.Println(c.Query("q"))
	response, err := http.Get("https://rule34.xxx/public/autocomplete.php?q=" + c.Query("q"))
	if err != nil {
		c.JSON(404, "Cannot fetch autocomplete")
		return
	}
	defer response.Body.Close()
	fmt.Println(response)
	fmt.Println(response.Body)

	autoc, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(404, "Cannot read autocomplete")
		return
	}

	fmt.Println(string(autoc))

	c.JSON(200, string(autoc))
}
