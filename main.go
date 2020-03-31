package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/gocarina/gocsv"

	log "github.com/sirupsen/logrus"
)

// WunderListFromFile reads a WunderlistExport from a JSON file
func WunderListFromFile(inputFile string) (WunderListExport, error) {
	jsonFile, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	byteValue = bytes.TrimPrefix(byteValue, []byte("\xef\xbb\xbf"))

	wunderlist := WunderListExport{}
	err = json.Unmarshal([]byte(byteValue), &wunderlist)
	if err != nil {
		return nil, err
	}
	return wunderlist, nil
}

// NotionListToFile writes a list of Notion tasks to the output file
func NotionListToFile(notionItems []*NotionExport, outputFile string) error {

	csvFile, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	err = gocsv.MarshalFile(notionItems, csvFile)
	if err != nil {
		return err
	}

	return nil

}

// WunderListTaskToNotionTask converts a single Wunderlist task object to a Notion task object
func WunderListTaskToNotionTask(in WunderListTask) (out *NotionExport) {
	out = &NotionExport{
		Name:        in.Title,
		DateCreated: in.CreatedAt,
		Status: func() string {
			if in.Completed {
				return "Done 🙌"
			}
			return "To Do"
		}(),
	}
	return
}

func main() {

	var (
		inputFile    = flag.String("input", "", "Input file in Wunderlist JSON format that will be parsed. (Required)")
		listToExport = flag.String("list", "inbox", "List to export.")
		outputFile   = flag.String("output", "", "Output file in Notion CSV that will be written. (Required)")
	)
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Parse Wunderlist JSON data
	log.Println("Opening input file")
	wunderlist, err := WunderListFromFile(*inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully parsed Wunderlist data")

	// Gather Wunderlist tasks to export
	var itemsToExport []WunderListTask

	for _, l := range wunderlist {
		if l.Title == *listToExport {
			itemsToExport = l.Tasks
		}
	}

	if itemsToExport == nil {
		log.Fatalf("Couldn't find list '%s'.", listToExport)
	}

	log.Printf("number of tasks to export: %v", len(itemsToExport))

	// Convert Wunderlist tasks to Notion tasks
	notionItems := []*NotionExport{}

	for _, item := range itemsToExport {
		notionItems = append(notionItems, WunderListTaskToNotionTask(item))
	}

	// Write Notion tasks to CSV
	err = NotionListToFile(notionItems, *outputFile)
	if err != nil {
		log.Fatalf("Couldn't save csv file %v.", err)
	}

	log.Printf("successfully saved tasks to Notion output %s", outputFile)

}
