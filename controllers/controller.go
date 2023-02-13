package controllers

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/ShikharY10/zintlr_internship_task_app_server/config"
	"github.com/ShikharY10/zintlr_internship_task_app_server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Controller struct {
	DB *config.DataBase
}

func (ctrl *Controller) TestRoute(c *gin.Context) {
	c.String(200, "Test Route is Working")
}

func (ctrl *Controller) RegisterUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("service", "Zintlr Task App")

	// collecting request body
	var request map[string]any
	c.BindJSON(&request)

	var username string = request["username"].(string)
	var tags []interface{} = request["tags"].([]interface{})
	var avatarData map[string]interface{} = request["avatar"].(map[string]interface{})

	avatar := models.Image{
		SecureUrl: avatarData["secureUrl"].(string),
		PublicId:  avatarData["publicId"].(string),
	}
	var tagss []string
	for _, element := range tags {
		tagss = append(tagss, element.(string))
	}

	user := models.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Tags:     tagss,
		Avatar:   avatar,
	}

	ctrl.DB.Users.InsertOne(context.TODO(), user)

	c.JSON(200, user)
}

func (ctrl *Controller) AddPost(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("service", "Zintlr Task App")

	// collecting request body
	var request map[string]any
	c.BindJSON(&request)

	var caption string = request["caption"].(string)

	var tags []string
	for _, element := range request["tags"].([]interface{}) {
		tags = append(tags, element.(string))
	}
	objectId := request["parentId"].(string)
	parentId, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		c.AbortWithStatusJSON(400, "Invalid ParentId")
	} else {
		var imageData string = request["image"].(string)

		if len(caption) > 0 && len(tags) > 0 && len(imageData) > 0 {
			post := models.Post{
				ID:                   primitive.NewObjectID(),
				ParentId:             parentId,
				ParentAvatarPublicId: request["parentAvatar"].(string),
				Caption:              caption,
				Tags:                 tags,
				ImagePublicId:        imageData,
			}
			ctrl.DB.Posts.InsertOne(context.TODO(), post)
			c.JSON(200, post)
		}
	}
}

func (ctrl *Controller) GetRandomPost(c *gin.Context) {
	id := c.Query("id")
	if len(id) > 0 {
		fmt.Println(`ID: ${id} | ` + id)
		findOptions := options.Find()
		findOptions.SetLimit(20)
		var results []models.Post

		cursor, err := ctrl.DB.Posts.Find(
			context.TODO(),
			bson.D{{}},
			findOptions,
		)
		if err != nil {
			c.AbortWithStatus(400)
		} else {
			for cursor.Next(context.TODO()) {
				var elem models.Post
				err := cursor.Decode(&elem)
				if err != nil {
					c.AbortWithStatus(400)
				} else {
					results = append(results, elem)
				}
			}

			if err := cursor.Err(); err != nil {
				c.AbortWithStatus(400)
			}
			cursor.Close(context.TODO())
			rand.Seed(time.Now().UnixNano())
			shuffle(results)
			c.JSON(200, results)
		}
	}
}

func (ctrl *Controller) DeleteAccount(c *gin.Context) {
	id := c.Query("id")
	if len(id) > 0 {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			fmt.Println("bad id")
			return
		}
		filter := bson.D{{Key: "_id", Value: objectId}}
		result, err := ctrl.DB.Users.DeleteOne(context.TODO(), filter)
		if err != nil {
			c.AbortWithStatusJSON(400, "error while deleting")
		} else {
			if result.DeletedCount > 0 {
				c.JSON(200, "Successfully Deleted")
			} else {
				c.AbortWithStatusJSON(400, "wrong")
			}
		}
	}
}

func shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		swap(i, j)
	}
}
