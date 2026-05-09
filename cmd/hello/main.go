package main

import (
  "flag"
  "github.com/PeterChen607/hellogo"
  "os"
)

var name string

func init() {
  flag.StringVar(&name, "name", "world", "name to say hello")
}

func main() {
  flag.Parse()
  msg := hellogo.Hello(name)
  _, err := os.Stdout.WriteString(msg)
  if err != nil {
    os.Stderr.WriteString(err.Error())
  }
}