package convert

import(
  "strconv"
  "errors"
  "strings"
  "path"
  "os/exec"
  "bytes"
)

func baseName(fileName string) (baseName string) {
  i, j := strings.LastIndex(fileName, "/") + 1, strings.LastIndex(fileName, path.Ext(fileName))
  return fileName[i:j]
}


func PdfToPng(f string, d int) (png string, err error) {
  var outputName = baseName(f) + ".png"
  cmd := exec.Command("convert", "-density", strconv.Itoa(d), f, outputName)
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    return "", err
  }

  if err := cmd.Start(); err != nil {
    return "", err
  }

  buf := new(bytes.Buffer)
  buf.ReadFrom(stdout)
  s := buf.String() // Does a complete copy of the bytes in the buffer.


  if err := cmd.Wait(); err != nil {
    return "", err
  } else {
    return outputName, nil
  }

  if s != "" {
    return "", errors.New("Could not convert: " + f)
  } else {
    return outputName, nil
  }
}
