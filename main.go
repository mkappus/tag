package main

import (
	"fmt"
	"log"

	"github.com/mkappus/tag/src/course"
)

func main() {
	c := course.Course{1, &course.Teacher{Id: 1}, "the thing", course.Students{}}
	cs, err := course.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := cs.Add(c); err != nil {
		log.Fatal(err)
	}
	fmt.Println("fini")

}
