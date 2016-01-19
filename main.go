package main

import(
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "os"
  "errors"
  "nycjv321.com/sort-a-bytes/latex"
  "nycjv321.com/sort-a-bytes/convert"
)

func randomDirectory() (randomDirectory string) {
  output, err := ioutil.TempDir(os.TempDir(), "latex-server")

  if err != nil {
    panic(err)
  } else {
    return output
  }
}

func createPdf(input []byte) (outputFile string, e error)  {
  var err = latex.CreatePdf(string(input))

  if err != nil {
    return "", err
  }

  if _, err = os.Stat("article.pdf"); os.IsNotExist(err) {
    return "", errors.New("Could not create: article.pdf")
  } else {
    return "article.pdf", nil
  }
}

func sanityCheck(r *gin.Engine) {
  r.GET("/sanity-check.json", func(c *gin.Context) {

    type tool struct {
      Name    string `json:"name"`
      Description string `json:"description"`
      Homepage  string `json:"homepage"`
      Installed bool `json:"installed"`
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
      "latex": sanity.latex,
      "pdflatex": sanity.pdflatex,
      "convert": sanity.convert,
    })

  })
}

func process(r *gin.Engine) {

    r.POST("/process", func(c *gin.Context) {
      f := c.DefaultQuery("format", "pdf")
      o, err := ioutil.ReadAll(c.Request.Body)
      if len(o) == 0 {
        c.String(400, "Empty data is not a valid form of input")
      } else if err != nil {
        c.String(400, err.Error())
      } else {
        if f == "pdf" || f == "png" {

          randomDirectory := randomDirectory()
          os.Chdir(randomDirectory)

          outputFile, err := createPdf(o)

          if err != nil {
            c.String(500, err.Error())
          }

          if f == "png" {
            var png, err = convert.PdfToPng(outputFile, 300)

            if err != nil {
              c.String(500, err.Error())
            }
            c.File(png)
            os.Remove(png)

          } else if f == "pdf" {
            c.File("article.pdf")
          }

          os.Remove(outputFile)
        } else {
          c.String(400, f + " is an unsupported format")
        }
      }
    })
}


func main() {
  r := gin.Default()
  r.LoadHTMLGlob("templates/*")
  r.Static("js", "build/js")
  r.Static("css", "build/css")

  process(r)
  sanityCheck(r)

  r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
