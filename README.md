Number Classification API

A Go-based API that takes a number as input and returns interesting mathematical properties about it, along with a fun fact.

Features

Checks if a number is prime or perfect.

Identifies Armstrong numbers.

Determines if a number is odd or even.

Computes the sum of its digits.

Retrieves a fun fact from the Numbers API.


API Specification

Endpoint:

GET /api/classify-number?number={number}

Query Parameter:

number (integer) – The number to classify.


Response Format:

Success Response (200 OK):

{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}

Error Response (400 Bad Request):

{
    "number": "alphabet",
    "error": true
}

Tech Stack

Language: Go (Golang)

Framework: Gorilla Mux

Deployment: Render

External API: Numbers API


Installation

1. Clone the repository:

git clone https://github.com/LazyShikamaru/Classification-.git
cd Classification-


2. Install dependencies:

go mod tidy


3. Run the API:

go run main.go



Deployment

The API is deployed on Render and can be accessed at:
🔗 Live API

Example Usage

curl -X GET "<your-deployment-url>/api/classify-number?number=371"

Contributing

Contributions are welcome! If you find a bug or have an improvement suggestion:

1. Fork the repository.


2. Create a new branch (feature-branch).


3. Commit your changes.


4. Push to your branch and open a pull request.



License

This project is open-source and available under the MIT License.


---

Let me know if you need any modifications!

