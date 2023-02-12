package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Tags     []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	Avatar   Image              `bson:"avatar,omitempty" json:"avatar,omitempty"`
}

type Image struct {
	SecureUrl string `bson:"secureUrl,omitempty" json:"secureUrl,omitempty"`
	PublicId  string `bson:"publicId,omitempty" json:"publicId"`
}

type Post struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ParentId             primitive.ObjectID `bson:"parentId,omitempty" json:"parentId,omitempty"`
	ParentAvatarPublicId string             `bson:"parentAvatar,omitempty" json:"parentAvatar,omitempty"`
	Caption              string             `bson:"caption,omitempty" json:"caption"`
	Tags                 []string           `bson:"tags,omitempty" json:"tags"`
	ImagePublicId        string             `bson:"image,omitempty" json:"image"`
}
