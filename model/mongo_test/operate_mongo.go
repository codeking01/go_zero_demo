package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/zeromicro/go-zero/core/stores/mon"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

func main() {
	conn := mon.MustNewModel("mongodb://root:123456@localhost:27017", "test_mongo_db", "collection")
	ctx := context.Background()

	u := &User{
		Username: "username",
		Password: "password",
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
	}

	// 插入文档
	result, err := conn.InsertOne(ctx, u)
	if err != nil {
		panic(err)
	}
	// 插入后的ID
	u.ID = result.InsertedID.(primitive.ObjectID)

	// 查询文档
	var newUser User
	err = conn.FindOne(ctx, &newUser, bson.M{"_id": u.ID})
	if err != nil {
		panic(err)
	}

	// 更新文档
	newUser.Username = "newUsername"
	_, err = conn.ReplaceOne(ctx, bson.M{"_id": newUser.ID}, newUser)
	if err != nil {
		panic(err)
	}

	// 删除文档
	_, err = conn.DeleteOne(ctx, bson.M{"_id": newUser.ID})
	if err != nil {
		panic(err)
	}
}
