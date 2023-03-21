package main

import (
	"fmt"

	"github.com/acoutinho1402/go-rest-api.git/models"
	"github.com/acoutinho1402/go-rest-api.git/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o Servidor Rest com GO")
	routes.HandleRequest()
}
