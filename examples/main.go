package main

import (
	"examples/a"
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	fmt.Printf("print from native\n")
	fmt.Printf("import package of sub-folder: a.A() => %s\n", a.A())
	j := simplejson.New()
	j.UnmarshalJSON([]byte("{\"id\": 1234567890}"))
	id, _ := j.Get("id").Int()
	fmt.Printf("import package of github.com/bitly/go-simplejson: {\"id\": %d}", id)
}
