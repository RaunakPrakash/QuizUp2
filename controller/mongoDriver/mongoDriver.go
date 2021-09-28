package mongoDriver

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rmq/controller/model"
)

func GetCollection(ctx context.Context,db , cl string) (*mongo.Collection,error){
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil {
		return nil, errors.Wrap(err,"Connection error in GetCollection Function")
	}
	collection := client.Database(db).Collection(cl)
	return collection,nil
}

func PutUser(ctx context.Context,collection *mongo.Collection,user model.Model) (interface{},error) {
	update := bson.M{"$set": bson.M{"level":user.Level,"total":user.Total,"points":user.Points,"date":user.Date}}
	res , err := collection.UpdateOne(ctx,bson.M{"username":user.Username},update)
	if err != nil {
		return nil,errors.Wrap(err,"Storing Problem")
	}
	fmt.Println(res.MatchedCount,res.UpsertedCount,res.ModifiedCount,res.UpsertedID)
	return res,nil
}

