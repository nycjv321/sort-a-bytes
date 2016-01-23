package main

import(
  "os/exec"
)

type SanityCheck struct {
    latex bool
    convert bool
    pdflatex bool
}

func GetSanityCheck() *SanityCheck {
  return &SanityCheck{latex: Installed("latex"), pdflatex: Installed("pdflatex"), convert: Installed("convert")}
}

func Installed(command string) bool {
  cmd := exec.Command("which", command)

  if err := cmd.Run(); err != nil {
   return false
  } else {
    return true
  }

}
