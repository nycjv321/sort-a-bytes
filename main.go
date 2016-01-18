package main

import(
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "os"
  "errors"
  "nycjv321.com/latex-server/latex"
  "nycjv321.com/latex-server/convert"
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
            c.String(501, err.Error())
          }

          if f == "png" {
            var png, err = convert.PdfToPng(outputFile, 300)

            if err != nil {
              c.String(502, err.Error())
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
  process(r)

  r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
