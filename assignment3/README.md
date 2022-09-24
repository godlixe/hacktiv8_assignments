## Hacktiv8 Assignment 3

### Description
The given task is to make a microservice that generates a random number from 1-100 for wind and water levels. the status will then be determined based on the levels of wind and water. The service will update a JSON file containing the data every 15 seconds. My approach to this problem is to make an API that has a specific service to run an updating goroutine.

Why did i make this project as an API? 
It's just for practicing, some of the class names are no-names, I tried following a common clean architecture approach. 

Future fixes :
- Make lock to ensure the file has content while fetching data from it
- Improve clean architecture

### Running the project
To run this project, clone this repository and  run `go run main.go` in the terminal.
The API can be accessed from port 8080

### Endpoints
| Endpoint                   | Description                          |
|----------------------------|--------------------------------------|
| GET http://localhost:8080/ | Retrieve current data from JSON file |
