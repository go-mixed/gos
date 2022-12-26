package main

import (
	"example1/a"
	"fmt"
	"github.com/bitly/go-simplejson"
	"os"
)

func main() {
	cd, _ := os.Getwd()
	exe, _ := os.Executable()
	fmt.Printf("---native---\n")
	fmt.Printf(" - os.Args: %#v\n", os.Args)
	fmt.Printf(" - working directory: %s\n", cd)
	fmt.Printf(" - application path: %s\n", exe)

	fmt.Printf("---inner package(sub-folder)---\n")
	fmt.Printf(" - a.A() => %s\n", a.A())

	fmt.Printf("---3rd package---\n")
	j := simplejson.New()
	j.UnmarshalJSON([]byte("{\"id\": 1234567890}"))
	id, _ := j.Get("id").Int()
	fmt.Printf(" - github.com/bitly/go-simplejson: simplejson.UnmarshalJSON => {\"id\": %d}\n", id)

}
