package main

import (
	"log"

	"github.com/rodkranz/test/module/instagram"
)

func main() {
	b, e := instagram.FetchFromTag("yoga")
	if e != nil {
		log.Fatalf("Error fetch from tag: %v", e)
	}

	b, e = instagram.ParseData(b)
	if e != nil {
		log.Fatalf("Error parse data: %v", e)
	}

	obj := new(instagram.Data)
	e = obj.Parser(b)
	if e != nil {
		log.Fatalf("Error parse obj: %v", e)
	}

	log.Printf("%v", obj)gs

}
