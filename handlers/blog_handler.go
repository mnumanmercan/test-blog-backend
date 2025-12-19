package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"blog-backend/database"
	"blog-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog
	json.NewDecoder(r.Body).Decode(&blog)

	collection := database.DB.Collection("posts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, _ := collection.InsertOne(ctx, blog)
	blog.ID = result.InsertedID.(interface{}).(primitive.ObjectID)

	json.NewEncoder(w).Encode(blog)
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var blogs []models.Blog
	collection := database.DB.Collection("posts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, _ := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var blog models.Blog
		cursor.Decode(&blog)
		blogs = append(blogs, blog)
	}

	json.NewEncoder(w).Encode(blogs)
}
