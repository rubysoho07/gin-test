package templates

import (
	"bytes"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type TemplateInput struct {
	Branch string `json:"branch" binding:"required"`
	Runner string `json:"runner" binding:"required"`
}

func GetFileFromTemplate(c *gin.Context) {

	var ti TemplateInput
	var r bytes.Buffer
	err := c.ShouldBindJSON(&ti)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tf, err := os.ReadFile("./ghaw_template.txt")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t := template.Must(template.New("template").Parse(string(tf)))

	err = t.Execute(&r, ti)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, r.String())
}
