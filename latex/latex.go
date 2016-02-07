package latex

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

func getLog() (log string) {
	entries, err := ioutil.ReadDir(".")
	if err != nil {
		return ""
	} else {
		for i := 0; i < len(entries); i++ {
			if strings.Contains(entries[i].Name(), ".log") {
				log, err := ioutil.ReadFile(entries[i].Name())
				if err != nil {
					return ""
				} else {
					return string(log)
				}
			}
		}
		return ""
	}
}

func CreatePdf(tex string) (path string, err error) {
	cmd := exec.Command("pdflatex")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stdin, err := cmd.StdinPipe()

	if err != nil {
		return "", err
	} else {
		io.WriteString(stdin, tex)
		stdin.Close()
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)
	s := buf.String() // Does a complete copy of the bytes in the buffer.

	if err := cmd.Wait(); err != nil {
		log := getLog()
		if len(log) > 0 {
			return "", errors.New(log)
		} else {
			return "", errors.New("Unable to compile .tex. Unable to read log")
		}
	} else {
		if !strings.Contains(s, "Output written") {
			return "", errors.New(s)
		} else {
			regex := regexp.MustCompile("[a-zA-Z]+[.]pdf")
			return regex.FindString(s), nil
		}
	}
}
