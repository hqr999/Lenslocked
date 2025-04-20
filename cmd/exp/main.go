package main

import (
	"os"
	"html/template"
)


type User struct {
		Nome string
		Bio string
}


func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
			panic(err)
	}

	u1 := User{
				Nome: "Ernst Jünger",
				Bio: `<script>alert("Haha você foi h4x0r3d!");</script>`,
	}

	err = t.Execute(os.Stdout,u1)

	if err != nil {
		panic(err)
	}

	
}
