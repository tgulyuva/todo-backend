package models

import "gopkg.in/mgo.v2/bson"

type Task struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	IsDone bool          `json:"IsDone" bson:"IsDone"`
}
