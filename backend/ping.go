package backend

import (
	"context"
	"net/http"
	"sync"
	"time"

	"gitlab.com/gaming0skar123/go/pingbot/common"
	"gitlab.com/gaming0skar123/go/pingbot/config"
)

var (
	checkErr   = common.CheckErr
	cacheRetry int
)

func ping() {
	if cacheRetry >= config.Toml.Backend.Cache {
		cache()
		cacheRetry = 0
	}
	cacheRetry++

	var wg sync.WaitGroup

	for _, url := range cacheURL {
		go loop(wg, url)
	}

	wg.Wait()
}

func loop(wg sync.WaitGroup, url string) {
	wg.Add(1)
	defer wg.Done()

	// Timeout 1 minute
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if checkErr(err, "new http request") {
		Status.Error++
		return
	}

	client := http.DefaultClient
	r, err := client.Do(req)
	if err != nil {
		Status.Error++
		return
	}

	if r.StatusCode >= 200 && r.StatusCode < 400 {
		Status.Success++
	} else {
		Status.Error++
	}
}
