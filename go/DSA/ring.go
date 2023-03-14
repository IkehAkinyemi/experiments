package main

import (
	"container/ring"
	"fmt"
)

func main (){
	//initialization 
  r := ring.New(2)
  s := ring.New(2)
  //populating 
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
		s.Value = i+1
		s = s.Next()
	}

	//Linking
	ls := r.Link(s)

	fmt.Println("linking r and s")
	ls.Do(func(p interface{}) {
			fmt.Println(p)
		})

	fmt.Println("Unlinking ls at point 3")
	ls.Unlink(3) //Unliking
	ls.Do(func(p interface{}) {
			fmt.Println(p)
		})

}