package main

import (
	"fmt"
	"log"
)

func main(){
	store , err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	serr := store.Init(); if serr != nil {
		log.Fatal(serr)
	}
	fmt.Printf("%+v\n",store)
	server := NewAPIServer(":3000",store)
	server.Run()
	fmt.Println("Hello wolrd")
}