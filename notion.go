package main

import "time"

// NotionExport represents the CSV export format of Notion
type NotionExport struct {
	Name        string    `csv:"Name"`
	DateCreated time.Time `csv:"Date Created"`
	Status      string    `csv:"Status"`
}
