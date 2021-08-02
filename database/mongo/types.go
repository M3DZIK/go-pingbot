package mongo

type URL struct {
	URL     string `bson:"_id"`
	Cluster int    `bson:"cluster"`
	Owner   string `bson:"owner"`
}

type json map[string]interface{}
