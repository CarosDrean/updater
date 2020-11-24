package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	log.Println("hello world!")
	files, err := ioutil.ReadDir("D:\\temp")
	if err != nil {
		log.Println(err)
	}

	for _, file := range files{
		fmt.Println("Nombre:", file.Name())
		fmt.Println("Tamaño:", file.Size())
		fmt.Println("Modo:", file.Mode())
		fmt.Println("Ultima modificación:", file.ModTime())
		fmt.Println("Es directorio?:", file.IsDir())
		fmt.Println("-----------------------------------------")
	}
}
