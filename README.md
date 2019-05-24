# Hotel_Automation

Pre-requistie:

Golang should be installed. Version > 1.10.1

Details:

$ make
------------------------------------------------------------------------
HOTEL AUTOMATION
------------------------------------------------------------------------
build                          build application binaries
lint                           check code for lint errors
test                           run unit tests
$ make lint
go vet ./...

$ make build
GOOS=darwin GOARCH=amd64 go build -o bin/hotelautomation-darwin-amd64 .
GOOS=linux GOARCH=amd64 go build -o bin/hotelautomation-linux-amd64 .
$ cd bin/
$ ls
hotelautomation-darwin-amd64	hotelautomation-linux-amd64
$ ./hotelautomation-darwin-amd64 
Hotel Automation. Please enter details:
Enter No. of Floors:
2
Enter No. of Main Corridor per Floor:
1
Enter No. of Sub Corridor per Floor:
2

Output from controller:

		Floor 1		
Main corridor 1 Light 1 : ON AC : ON
Sub corridor 1 Light 1 : OFF AC : ON
Sub corridor 2 Light 2 : OFF AC : ON
		Floor 2		
Main corridor 1 Light 1 : ON AC : ON
Sub corridor 1 Light 1 : OFF AC : ON
Sub corridor 2 Light 2 : OFF AC : ON

Checking Sensor Inputs...
Movement in Floor 1, Sub corridor 1

Output from controller:

		Floor 1		
Main corridor 1 Light 1 : ON AC : ON
Sub corridor 1 Light 1 : ON AC : ON
Sub corridor 2 Light 2 : OFF AC : OFF
		Floor 2		
Main corridor 1 Light 1 : ON AC : ON
Sub corridor 1 Light 1 : OFF AC : ON
Sub corridor 2 Light 2 : OFF AC : ON

Checking Sensor Inputs...
No Movement in Floor 1, Sub corridor 1 for a minute

Output from controller:

		Floor 1		
Main corridor 1 Light 1 : ON AC : ON
Sub corridor 1 Light 1 : OFF AC : ON
Sub corridor 2 Light 2 : OFF AC : ON
		Floor 2		
Main corridor 1 Light 1 : ON AC : ON
Sub corridor 1 Light 1 : OFF AC : ON
Sub corridor 2 Light 2 : OFF AC : ON

Checking Sensor Inputs...
^C
$ 


