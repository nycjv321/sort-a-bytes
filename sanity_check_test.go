package main

import "testing"

func TestInstalled(t *testing.T) {
  if !Installed("pdflatex") {
    t.Fatal("pdflatex isn't installed")
  }

}

func TestNotInstalled(t *testing.T) {
  if Installed("qwerty!!") {
    t.Fatal("Invalid command should not be considered as valid")
  }
}
