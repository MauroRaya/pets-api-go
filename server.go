package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pet struct {
	Id          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Species     string `json:"species" binding:"required"`
	DateOfBirth string `json:"dateOfBirth" binding:"required"`
}

var pets = []pet{
	{Id: "1", Name: "Flakes", Species: "Dog", DateOfBirth: "10/05/2015"},
	{Id: "2", Name: "Peppy", Species: "Parrot", DateOfBirth: "20/06/2004"},
	{Id: "3", Name: "Jake", Species: "Cat", DateOfBirth: "12/01/2019"},
}

func getPetById(id string) (*pet, error) {
	for i := range pets {
		if pets[i].Id == id {
			return &pets[i], nil
		}
	}
	return nil, errors.New("pet not found")
}

func findPetIndex(id string) (int, error) {
	for i := range pets {
		if pets[i].Id == id {
			return i, nil
		}
	}
	return -1, errors.New("pet not found")
}

func getPets(context *gin.Context) {
	context.JSON(http.StatusOK, pets)
}

func getPet(context *gin.Context) {
	id := context.Param("id")

	pet, err := getPetById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, pet)
}

func addPet(context *gin.Context) {
	var newPet pet

	if err := context.BindJSON(&newPet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, existingPet := range pets {
		if existingPet.Id == newPet.Id {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "id already in use",
			})
			return
		}
	}

	pets = append(pets, newPet)

	context.JSON(http.StatusCreated, newPet)
}

func editPet(context *gin.Context) {
	id := context.Param("id")

	i, err := findPetIndex(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var editedPet pet

	if err := context.BindJSON(&editedPet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	pets[i].Name = editedPet.Name
	pets[i].Species = editedPet.Species
	pets[i].DateOfBirth = editedPet.DateOfBirth

	context.JSON(http.StatusOK, gin.H{
		"message": "Pet edited successfully",
		"pet":     pets[i],
	})
}

func removePet(context *gin.Context) {
	id := context.Param("id")

	i, err := findPetIndex(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	deletedPet := pets[i]

	pets = append(pets[:i], pets[i+1:]...)

	context.JSON(http.StatusOK, gin.H{
		"message": "Pet removed successfully",
		"pet":     deletedPet,
	})
}

func main() {
	router := gin.Default()

	router.GET("/pets", getPets)
	router.GET("/pets/:id", getPet)
	router.POST("/pets", addPet)
	router.PATCH("/pets/:id", editPet)
	router.DELETE("/pets/:id", removePet)

	router.Run("0.0.0.0:80")
}
