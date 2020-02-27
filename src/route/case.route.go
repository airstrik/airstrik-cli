package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetCase(c *gin.Context) {
	fmt.Printf("%s", "At least I got here")
	c.JSON(200, gin.H{"status": "message"}) //???
}

func AddOrUpdateCase(c *gin.Context) {

}

func DeleteCase(c *gin.Context) {
	fmt.Printf("%s", "At least I got here")
	c.JSON(200, gin.H{"status": "message"}) //???
}
