package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "OK"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
