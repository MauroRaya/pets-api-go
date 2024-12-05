package main

import (
	"github.com/gin-gonic/gin"
	"errors"
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

func getPetById(id string) (*pet, error) {
	for i := range pets {
		if pets[i].Id == id {
			return &pets[i], nil
		}
	}
	return nil, errors.New("pet not found")
}

func getPets(context *gin.Context) {
	context.JSON(http.StatusOK, pets)
}

func getPet(context *gin.Context) {
	id := context.Param("id")

	pet, err := getPetById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, pet)
}

func main() {
	router := gin.Default()

	router.GET("/pets",     getPets)
	router.GET("/pets/:id", getPet)	

	router.Run("localhost:80")
}
