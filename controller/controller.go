package controller

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"rmq/controller/model"
	"rmq/controller/mongoDriver"
)

type UserController struct {
	collection *mongo.Collection
}

func (this *UserController) SetCollection(ctx context.Context,db,cl string) error {
	collection, err := mongoDriver.GetCollection(ctx,db,cl)
	if err != nil {
		return err
	}
	this.collection = collection
	return nil
}

func (this *UserController) Put(ctx context.Context,user model.Model) (interface{},error) {

	id , err := mongoDriver.PutUser(ctx,this.collection,user)
	if err != nil {
		return nil,errors.Wrap(err,"Put problem")
	}
	return id,nil
}
