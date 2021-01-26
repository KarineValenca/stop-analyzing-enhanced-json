# STOP ANALYZING ENHANCED JSON
This project generates an enhanced JSON file to be used in the Stop Analyzing project. It reads a JSON file called `lafiancee.json` and makes a request to the lafiancee store webpage. Finally, it creates a file called `enhanced.json` with the JSON that will be used by the Stop Analyzing project.

## How-To
Clone this project and run:

`go run main.go`

If you need to get more products from the lafiancee store, just update the `lafiancee.json` file with the new products. Please, keep the same JSON data structure. Then, run the program again, and you'll get the enhanced JSON updated with the new data in the `enhanced.json` file. 