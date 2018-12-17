package models

import "gopkg.in/mgo.v2/bson"

// Using 'bson' kwyworkd to tell mgo drive how to name properties in mongo db docoment
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}
