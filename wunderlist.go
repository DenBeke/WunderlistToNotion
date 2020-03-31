package main

import "time"

// WunderListExport represents the JSON format of a Wunderlist export
type WunderListExport []WunderListList

type WunderListList struct {
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	Folder        interface{}      `json:"folder"`
	Tasks         []WunderListTask `json:"tasks"`
	DirectoryPath string           `json:"directoryPath"`
	TextFilePath  string           `json:"textFilePath"`
	HTMLFilePath  string           `json:"htmlFilePath"`
	CreatedAt     time.Time        `json:"createdAt"`
}

type WunderListTask struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Completed bool          `json:"completed"`
	Starred   bool          `json:"starred"`
	Subtasks  []interface{} `json:"subtasks"`
	Notes     []interface{} `json:"notes"`
	Comments  []interface{} `json:"comments"`
	Reminders []interface{} `json:"reminders"`
	Files     []interface{} `json:"files"`
	Assignee  interface{}   `json:"assignee"`
	DueDate   interface{}   `json:"dueDate"`
	CreatedAt time.Time     `json:"createdAt"`
	CreatedBy struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"createdBy"`
	CompletedAt time.Time `json:"completedAt"`
	CompletedBy struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"completedBy"`
}
