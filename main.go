package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"nycjv321.com/sort-a-bytes/convert"
	"nycjv321.com/sort-a-bytes/latex"
)

func randomDirectory() (randomDirectory string) {
	output, err := ioutil.TempDir(os.TempDir(), "latex-server")

	if err != nil {
		panic(err)
	} else {
		return output
	}
}

func createPdf(input []byte) (outputFile string, e error) {
	var f, err = latex.CreatePdf(string(input))

	if err != nil {
		return "", err
	}

	if _, err = os.Stat(f); os.IsNotExist(err) {
		return "", errors.New("Could not create: output pdf")
	} else {
		return f, nil
	}
}

func sanityCheck(r *gin.Engine) {
	r.GET("/sanity-check.json", func(c *gin.Context) {

		type tool struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Homepage    string `json:"homepage"`
			Installed   bool   `json:"installed"`
		}

		var sanity = *GetSanityCheck()
		c.JSON(200, gin.H{
			"tools": []tool{
				tool{Name: "LaTeX", Description: "Compiles Tex", Homepage: "https://www.latex-project.org", Installed: sanity.latex},
				tool{Name: "pdflatex", Description: "Converts .tex to .pdf", Homepage: "https://www.latex-project.org", Installed: sanity.pdflatex},
				tool{Name: "convert", Description: "Converts .pdf to .png", Homepage: "http://www.imagemagick.org/script/index.php", Installed: sanity.convert},
			},
		})
	})
	r.GET("/sanity-check", func(c *gin.Context) {
		var sanity = *GetSanityCheck()
		c.HTML(200, "sanity-check.tmpl", gin.H{
			"latex":    sanity.latex,
			"pdflatex": sanity.pdflatex,
			"convert":  sanity.convert,
		})

	})
}

type Message struct {
	Value string `json:"value"`
}

func process(r *gin.Engine) {

	r.POST("/process", func(c *gin.Context) {
		f := c.DefaultQuery("format", "pdf")
		o, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, Message{Value: err.Error()})
		} else if len(o) == 0 {
			c.JSON(400, Message{Value: "Empty Input is Not Valid Form Submission"})
		} else {
			if f == "pdf" || f == "png" {
				randomDirectory := randomDirectory()
				os.Chdir(randomDirectory)

				outputFile, err := createPdf(o)

				if err != nil {
					c.JSON(500, Message{Value: err.Error()})
				} else {
					if f == "png" {
						var png, err = convert.PdfToPng(outputFile, 300)
						if err != nil {
							c.String(500, err.Error())
						}
						c.File(png)
						os.Remove(png)

					} else if f == "pdf" {
						c.File(outputFile)
					}
					os.Remove(outputFile)
				}
			} else {
				msg := Message{Value: fmt.Sprintf("\"%v\" was an invalid format", f)}
				c.JSON(400, msg)
			}
		}
	})
}

func Gin() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("js", "build/js")
	r.Static("css", "build/css")

	process(r)
	sanityCheck(r)
	return r
}

func main() {
	Gin().Run(":8080") // listen and serve on 0.0.0.0:8080
}
