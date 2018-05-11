package elasticsearch

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

type ESQueryOptions struct {
	Host  string
	Index string
	Types []string
	Sort  string
	Size  int
}

type Document struct {
	ID     string           `json:"_id"`
	Type   string           `json:"_type"`
	Source *json.RawMessage `json:"_source"`
}

func Query(opts *ESQueryOptions) ([]byte, error) {
	client, err := elastic.NewClient(elastic.SetURL(opts.Host))
	if err != nil {
		return nil, err
	}

	searchResult, err := client.Search().
		Index(opts.Index).
		Type(opts.Types...).
		Sort(opts.Sort, true).
		From(0).Size(opts.Size).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	var docs []Document
	for _, d := range searchResult.Hits.Hits {
		d := Document{d.Id, d.Type, d.Source}
		docs = append(docs, d)
	}

	logrus.WithField("amount", len(docs)).Info("documents found")

	b, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		return nil, err
	}

	return b, nil
}
