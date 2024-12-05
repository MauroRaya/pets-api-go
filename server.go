package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type pet struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Species     string `json:"species"`
	DateOfBirth string `json:"dateOfBirth"`
}

var pets = []pet {
	{Id: "1", Name: "Flakes", Species: "Dog",    DateOfBirth: "10/05/2015"},
	{Id: "2", Name: "Peppy",  Species: "Parrot", DateOfBirth: "20/06/2004"},
	{Id: "3", Name: "Jake",   Species: "Cat",    DateOfBirth: "12/01/2019"},
}

func getPets(context *gin.Context) {
	context.JSON(http.StatusOK, pets)
}

func main() {
	router := gin.Default()

	router.GET("/pets", getPets)	

	router.Run("localhost:80")
}
