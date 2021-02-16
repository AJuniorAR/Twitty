package main

import (
	"log"

	"github.com/AJuniorAR/Twitty/bd"
	"github.com/AJuniorAR/Twitty/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
