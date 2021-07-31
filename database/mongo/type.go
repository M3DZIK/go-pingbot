package mongo

type URL struct {
	URL     string `bson:"_id"`
	Cluster int    `bson:"cluster"`
}

type json map[string]interface{}
