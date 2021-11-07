package ping

import (
	"context"
	"net/http"
	"time"

	"github.com/MedzikUser/go-utils/common"
	"gitlab.com/MedzikUser/go/pingbot/config"
)

func ping() int {
	Cache()

	for _, url := range CacheURL {
		go loop(url)
	}

	return len(CacheURL)
}

func loop(url string) {
	// Timeout 1 minute
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		common.Log.Error("new http request", err)

		Status.Error++

		return
	}

	req.Header.Set("User-Agent", config.UserAgent)

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
