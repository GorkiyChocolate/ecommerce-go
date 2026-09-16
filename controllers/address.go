package controllers

import (
	"context"
	"ecommerce-go/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid code"})
			c.Abort()
			return
		}

		address, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid user id")
			return
		}

		var addresses models.Address

		addresses.Address_id = primitive.NewObjectID()

		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		match_filter := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: address}}}}

		unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$address"}}}}

		group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$address"}, {Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}}}}}

		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter, unwind, group})

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		var addressinfo []bson.M
		if err = pointcursor.All(ctx, &addressinfo); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		var size int32

		for _, addressNo := range addressinfo {
			if count, ok := addressNo["count"].(int32); ok {
				size = count
			}
		}
		if size < 2 {
			filter := bson.D{{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{{Key: "address", Value: addresses}}}}

			_, err := UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err.Error())
				return
			}
			c.IndentedJSON(http.StatusOK, "Address added successfully")
		} else {
			c.IndentedJSON(http.StatusBadRequest, "Not Allowed")
		}

		defer cancel()
	}
}

func EditHomeAddress() gin.HandlerFunc {

}

func EditWorkAddress() gin.HandlerFunc {

}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid search index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)

		userid, err := primitive.ObjectIDFromHex(user_id)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Invalid user id")
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		filter := bson.D{{Key: "_id", Value: userid}}

		update := bson.D{{Key: "$set", Value: bson.D{{Key: "address", Value: addresses}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)

		if err != nil {
			c.IndentedJSON(404, "wrong command")
			return
		}

		c.IndentedJSON(200, "Successfully deleted")

	}
}
