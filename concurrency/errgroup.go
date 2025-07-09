package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

/*
	errgroup - один из способов синхронизации горутин. Это синтаксический сахар, удобный в некоторых случаях
*/

var urls = []string{"https://example.com", "https://example.org", "https://example.net"}

func isAvailable(ctx context.Context, url string) error {
	c := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "OPTIONS", url, nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wrong status code %d for url %s", resp.StatusCode, url)
	}
	return nil
}
func main() {
	ctx := context.Background()
	g, gctx := errgroup.WithContext(ctx)
	for _, url := range urls {
		g.Go(func() error {
			return isAvailable(qctx, url)
		})
	}
	if err := g.Wait(); err != nil {
		log.Fatal("Some resource is not available: %v", err)
	} else {
		fmt.Println("All resources is ok.")
	}
}
