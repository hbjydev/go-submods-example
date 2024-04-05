package main

import (
  "github.com/hbjydev/go-submods-example"
  "github.com/hbjydev/go-submods-example/submod"
)

func main() {
  gosubmodsexample.Hello()
  submod.HelloWithName("hayden")
}
