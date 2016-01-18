package latex

import(
  "os/exec"
  "strings"
  "bytes"
  "io"
  "errors"
)

func CreatePdf(tex string) (err error) {

  cmd := exec.Command("pdflatex")
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    return err
  }

  stdin, err := cmd.StdinPipe()

  if err != nil {
    return err
  } else {
    io.WriteString(stdin, tex)
    stdin.Close()
  }

  if err := cmd.Start(); err != nil {
    return err
  }

  buf := new(bytes.Buffer)
  buf.ReadFrom(stdout)
  s := buf.String() // Does a complete copy of the bytes in the buffer.


  if err := cmd.Wait(); err != nil {
    return err
  } else {
    if !strings.Contains(s, "Output written on article.pdf") {
      return errors.New(s)
    }
  }
  return nil
}
