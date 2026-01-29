// Package rss contains all related rss functionality
// for gator
package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed

	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, err
	}

	newFeed := unescapeString(&feed)

	return newFeed, nil
}

func unescapeString(data *RSSFeed) *RSSFeed {
	data.Channel.Title = html.UnescapeString(data.Channel.Title)
	data.Channel.Description = html.UnescapeString(data.Channel.Description)
	for i := range data.Channel.Item {
		data.Channel.Item[i].Title = html.UnescapeString(data.Channel.Item[i].Title)
		data.Channel.Item[i].Description = html.UnescapeString(data.Channel.Item[i].Description)
	}

	return data
}
