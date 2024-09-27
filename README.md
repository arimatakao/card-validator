<div align="center">

# card validator 💳

Simple REST API for validating bank cards

</div>

# Features ⭐

- Check date (month and year);
- Validate card number using the Luhn algorithm;
- Logging;
- Graceful shutdown;
- Unit tests for validation func.

# Installation ⚙️

1. Clone this repository:

```sh
git clone https://github.com/arimatakao/card-validator.git
```

2. Navigate to the card-validator directory and build the Docker image:

```sh
docker build -t card-validator:latest .
```

3. Run the Docker image:

```sh
docker run -p 8080:8080 card-validator:latest
```

Log messages will print on your terminal:

```
2024/09/27 12:34:56 Server started and listeting on port :8080
...
```

You can run it natively:

```sh
make run
# or
go run main.go
```

For shutdown server send `SIGINT` signal (press `Ctrl + C`) in terminal and get log messages:

```
...
2024/09/27 12:35:11 Shutdown server
2024/09/27 21:35:14 Shutdown is successful
```

# Run tests 🚩

Natively:

```sh
make test
#or
go test -v ./...
```

Run tests in docker:

```sh
docker build -f Dockerfile-test -t card-validator-test:latest .
docker run card-validator-test
```

# Usage examples 📦

Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json' \
  --data '{
	"card_number": "5390802565179716",
	"expiration_month": "10",
	"expiration_year": "2030"
}'
```

Response:

```json
{
	"valid": true
}
```

---


Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json' \
  --data '{
	"card_number": "53908025651",
	"expiration_month": "10",
	"expiration_year": "2030"
}'
```

Response:

```json
{
	"valid": false,
	"error": {
		"code": 1,
		"message": "Card number must be between 12 and 19 digits"
	}
}
```

---
Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json' \
  --data '{
	"card_number": "5390802565179abc",
	"expiration_month": "10",
	"expiration_year": "2030"
}'
```

Response:

```json
{
	"valid": false,
	"error": {
		"code": 2,
		"message": "Card number contains invalid characters"
	}
}
```

---

Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json' \
  --data '{
	"card_number": "5390802565179715",
	"expiration_month": "10",
	"expiration_year": "2030"
}'
```

Response:

```json
{
	"valid": false,
	"error": {
		"code": 3,
		"message": "Card number is not valid"
	}
}
```

---

Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json' \
  --data '{
	"card_number": "5390802565179716",
	"expiration_month": "",
	"expiration_year": "2030"
}'
```

Response:

```json
{
	"valid": false,
	"error": {
		"code": 4,
		"message": "Expiration month field is required"
	}
}
```

---

Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json' \
  --data '{
	"card_number": "5390802565179716",
	"expiration_month": "abc",
	"expiration_year": "2030"
}'
```

Response:

```json
{
	"valid": false,
	"error": {
		"code": 5,
		"message": "Expiration month must be a number"
	}
}
```

---

Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json' \
  --data '{
	"card_number": "5390802565179716",
	"expiration_month": "10",
	"expiration_year": ""
}'
```

Response:

```json
{
	"valid": false,
	"error": {
		"code": 6,
		"message": "Expiration year field is required"
	}
}
```

---

Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json' \
  --data '{
	"card_number": "5390802565179716",
	"expiration_month": "10",
	"expiration_year": "abc"
}'
```

Response:

```json
{
	"valid": false,
	"error": {
		"code": 7,
		"message": "Expiration year must be a number"
	}
}
```

---

Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json' \
  --data '{
	"card_number": "5390802565179716",
	"expiration_month": "10",
	"expiration_year": "2011"
}'
```

Response:

```json
{
	"valid": false,
	"error": {
		"code": 8,
		"message": "The card has expired"
	}
}
```

---

Request:

```sh
curl --request POST \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json'
```

Response is empty body with 400 http status code.

---

Request:

```sh
curl --request GET \
  --url http://localhost:8080/api/validation \
  --header 'Content-Type: application/json'
```

Response is empty body with 405 http status code.
