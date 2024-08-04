package controllers

import (
	"golang-gin-db-restapi-lokal/database"
	"golang-gin-db-restapi-lokal/entities"
	"golang-gin-db-restapi-lokal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllAfkeur(c *gin.Context) {
	var result gin.H

	person, err := repository.GetAllAfkeur(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": person,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InserAfkeur(c *gin.Context) {
	var afkeur entities.Afkeur

	err := c.BindJSON(&afkeur)
	if err != nil {
		panic(err)
	}

	err = repository.InserAfkeur(database.DbConnection, afkeur)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, afkeur)
}

func UpdateAfkeur(c *gin.Context) {
	var afkeur entities.Afkeur
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&afkeur)
	if err != nil {
		panic(err)
	}

	afkeur.ID = id

	err = repository.UpdateAfkeur(database.DbConnection, afkeur)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, afkeur)
}

func DeleteAfkeur(c *gin.Context) {
	var afkeur entities.Afkeur
	// id, _ := strconv.Atoi(c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	afkeur.ID = id

	// err := repository.DeleteAfkeur(database.DbConnection, afkeur)
	err = repository.DeleteAfkeur(database.DbConnection, afkeur)
	if err != nil {
		// panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// c.JSON(http.StatusOK, afkeur)
	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully", "id": id})

}
