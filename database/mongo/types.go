package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type URL struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	URL     string             `bson:"url,omitempty" json:"url,omitempty"`
	Cluster int                `bson:"cluster,omitempty" json:"cluster,omitempty"`
}
