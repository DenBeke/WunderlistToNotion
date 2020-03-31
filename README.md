# Wunderlist to Notion converter

Converts Wunderlist JSON dump to a CSV file that can be imported into Notion.

⚠️ This is a work in progress. Backup the files you run it on and run it on your own risk. ⚠️

## Usage

    go run *.go --input=wunderlist_dump/Tasks.json --output=file_for_notion.csv


By default it exports the tasks from the Wunderlist 'inbox'. If you want to export another list, you can specify it with `--list=my-list`.

## Author

[Mathias Beke](https://denbeke.be)