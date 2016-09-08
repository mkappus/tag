package main

import (
	"fmt"
	"log"

	"github.com/mkappus/tag/src/course"
)

func main() {
	s := course.Students{
		{1, 7, "GU1", "CHRRISa", "ChrisanthIUM"},
		{2, 8, "GU2", "MARZZA,", "SHIVV"},
	}

	t := &course.Teacher{10, "Mistta", "Myzzah", "mm@aps.edu", "pwwww"}

	c := course.Course{2, t, "courrrses", s}
	cs, err := course.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := cs.Create(c); err != nil {
		log.Fatal(err)
	}
	fmt.Println("fini")

}
