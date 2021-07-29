package redis

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

type StatusType struct {
	Success int64
	Error   int64
}

func StatusUpdate(status StatusType) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := "status" + strconv.Itoa(config.Toml.Cluster.ID) + " " + strconv.Itoa(config.Toml.Cluster.Node)

	b, err := json.Marshal(status)
	if common.CheckErr(err, "redis: marshal json") {
		return
	}

	DB.Set(ctx, key, string(b), 0)
}

func StatusGet() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := "status" + strconv.Itoa(config.Toml.Cluster.ID) + " " + strconv.Itoa(config.Toml.Cluster.Node)

	s := DB.Get(ctx, key)
	println(s.Val())
}
