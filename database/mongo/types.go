package mongo

type URL struct {
	URL     string `bson:"_id,omitempty" json:"url,omitempty"`
	Cluster int    `bson:"cluster,omitempty" json:"cluster,omitempty"`
}
