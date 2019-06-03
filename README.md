# scooters
Ridecell Scooters API Project

# Configuration
.env file has all the config parameters.
By default the server is listening on port 8240.
To change the port, modify 'PORT' in the .env.

PostGreSQL configuration is also in .env file.
By default 'scooters' database is created with following tables:
- scooters
- reservations
- payments

# Pre-requistes
- PostGreSQL
- go get -u github.com/gorilla/mux
- go get -u github.com/jinzhu/gorm
- go get -u github.com/joho/godotenv

# Build and Deploy
- git clone https://github.com/murtigujar/scooters
- go build cmd/main.go
- go run cmd/main.go

# APIs
localhost:8420/api/v1/scooters
- GET: List all scooters

localhost:8420/api/v1/scooters?lat=17.788989&lng=-12.404810
- POST: Add new scooter

localhost:8420/api/v1/scooters/available?lat=17.78&lng=-12.40&radius=10
- GET: List all scooters with radius of given co-ordinates

localhost:8420/api/v1/scooters/reservations
- GET: List all reservations

localhost:8420/api/v1/scooters/reserve?id=2
- POST: Reserve scooter with id=2

localhost:8420/api/v1/scooters/return?res_id=2
- POST: End reservation for reservation id=2. This will mark the associated scooter available. This will also create a payment entry with the amount to be paid.

localhost:8420/api/v1/scooters/payments
- GET: List all payments

localhost:8420/api/v1/scooters/pay?res_id=2
- POST: Mock payment for the reservation id=2, by marking the amount as paid.