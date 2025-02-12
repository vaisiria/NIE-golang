package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

type Car struct {
	Brand    string `json:"brand"`
    Number   string `json:"number"`
    Type     string `json:"type"`
    Incoming string `json:"incoming"`
    Outgoing string `json:"outgoing"`
    Slot     string `json:"slot"`
    ID       string `json:"id"`
}

var dataFile = "data.json"

func loadCars() ([]Car, error) {
    file, err := os.Open(dataFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var cars []Car
    err = json.NewDecoder(file).Decode(&cars)
    if err != nil {
        return nil, err
    }
    return cars, nil
}

func saveCars(cars []Car) error {
    file, err := json.MarshalIndent(cars, "", " ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(dataFile, file, 0644)
}

func createCar(c *gin.Context) {
    var newCar Car
    if err := c.ShouldBindJSON(&newCar); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    cars, _ := loadCars()
    cars = append(cars, newCar)
    saveCars(cars)

    c.JSON(http.StatusCreated, gin.H{
        "message": "Car parking created successfully!",
        "car":     newCar,
    })
}

func getCars(c *gin.Context) {
    cars, err := loadCars()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading car parking details"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "Car parking details fetched successfully!",
        "car":     cars,
    })
}

func getCar(c *gin.Context) {
    id := c.Param("id")
    cars, _ := loadCars()

    for _, car := range cars {
        if car.ID == id {
            c.JSON(http.StatusOK, gin.H{
                "message": "Car parking details fetched successfully!",
                "car":     car,
            })
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Car details not found"})
}

func updateCar(c *gin.Context) {
    id := c.Param("id")
    var updatedCar Car
    if err := c.ShouldBindJSON(&updatedCar); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Invalid request payload"})
        return
    }
    cars, _ := loadCars()
    for i, car := range cars {
        if car.ID == id {
            cars[i] = updatedCar
            saveCars(cars)
            c.JSON(http.StatusOK, gin.H{
                "message": "Car parking details updated successfully!",
                "car":     updatedCar,
            })
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Car details not found"})
}

func deleteCar(c *gin.Context) {
    id := c.Param("id")
    cars, _ := loadCars()

    for i, car := range cars {
        if car.ID == id {
            cars = append(cars[:i], cars[i+1:]...)
            saveCars(cars)
            c.JSON(http.StatusOK, gin.H{
                "message": "Car deleted successfully!",
            })
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Car details not found"})
}

func main() {
    r := gin.Default()

    r.GET("/cars", getCars)
    r.GET("/cars/:id", getCar)
    r.POST("/cars", createCar)
    r.PUT("/cars/:id", updateCar)
    r.DELETE("/cars/:id", deleteCar)

    port := "8080"
    fmt.Printf("Server is running at http://localhost:%s\n", port)
    r.Run(":" + port)
}