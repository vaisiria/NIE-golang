package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Car struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Brand    string 			`json:"brand" bson:"brand"`
	Number   string 			`json:"number" bson:"number"`
	Type     string 			`json:"type" bson:"type"`
	Incoming string 			`json:"incoming_time" bson:"incoming_time"`
	Outgoing string 			`json:"outgoing_time" bson:"outgoing_time"`
	Slot     string 			`json:"parking_slot" bson:"parking_slot"`
}

var collection *mongo.Collection

func initMongoDB(){
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database("carparking").Collection("cars")
}

func createCar(c *gin.Context) {
	var newCar Car
	if err := c.ShouldBindJSON(&newCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newCar.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(context.TODO(), newCar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to insert the car parking details"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "Car added successfully!",
		"car":     newCar,
	})
}

func getCars(c *gin.Context) {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching car parking details"})
		return
	}
	defer cursor.Close(context.TODO())

	var cars []Car
	for cursor.Next(context.TODO()) {
		var car Car
		if err := cursor.Decode(&car); err != nil{
			continue
		}
		cars = append(cars, car)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cars retrieved successfully",
		"cars":    cars,
	})
}

func getCar(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	
	var car Car
	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&car)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Car retrieved successfully",
		"car":     car,
	})
}

func updateCar(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	
	var updatedCar Car
	if err := c.ShouldBindJSON(&updatedCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": updatedCar})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Car updated successfully!",
		"car":     updatedCar,
	})
}

func deleteCar(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the car parking details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully!"})	
}

func main() {
	initMongoDB()
	r := gin.Default()

	// Enable CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Anyone can access API
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173"}, // React Frontend URL
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Content-Type"},
	// 	AllowCredentials: true,
	// }))

	r.GET("/cars", getCars)
	r.GET("/cars/:id", getCar)
	r.POST("/cars", createCar)
	r.PUT("/cars/:id", updateCar)
	r.DELETE("/cars/:id", deleteCar)

	port := "8080"
	fmt.Printf("Server is running at: http://localhost:%s\n", port)
	r.Run(":" + port)
}