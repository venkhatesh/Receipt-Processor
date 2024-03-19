# Receipt Processor Service

## Overview

Receipt Processor is a service designed to process receipt data, calculate reward points based on specific rules, and manage receipt information in memory. This service provides an API to submit receipts and query calculated points.

## Features

- Process receipt JSON payloads.
- Calculate points based on the total amount, purchase time, item details, and other criteria.
- Store receipt data in memory.
- Retrieve calculated points by receipt ID.

## Getting Started

These instructions will get your copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.x or later)
- [Docker](https://docs.docker.com/get-docker/) (for containerization)

### Installing

A step-by-step series of examples that tell you how to get a development environment running.

1. **Clone the Repository**

   ```sh
   git clone https://github.com/yourusername/receipt-processor.git
   cd receipt-processor
   ```
2. **Using Docker**
   Build the Docker Image:

   ```sh
   docker build -t receipt-processor .
   ```
   Run the Docker container:

      ```sh
    docker run -p 8080:8080 receipt-processor
      ```

3. **Build the project (if you have GO installed)**

   ```sh
   go build .
   ```

   Run the service
   
   ```sh
   go run .
   ```

   Or, if you built the project
   
      ```sh
      ./Receipt-Processor
      ```
## API Reference

The service exposes two main endpoints:

### Process Receipts

- **Endpoint:** `/receipts/process`
- **Method:** POST
- **Payload:** Receipt JSON
- **Description:** Accepts a JSON receipt and returns an ID.

### Get Points

- **Endpoint:** `/receipts/{id}/points`
- **Method:** GET
- **Response:** A JSON object containing the number of points awarded.

## Running the Tests

To run the automated tests for this project, use the following command in your terminal:

```sh
go test ./...
```

This command will recursively run all tests in the current directory and subdirectories.

### Test Cases Covered

- **ProcessReceiptsEndpoint:** Tests the ```/receipts/process``` endpoint to ensure it correctly processes the input JSON and returns a unique ID.

- **GetPointsEndpoint:** Tests the ```/receipts/{id}/points``` endpoint to verify it returns the correct number of points based on the receipt data.

- **CalculatePointsLogic:** Validates the logic for calculating points from a receipt, ensuring points are accurately awarded based on the defined rules.

- **ErrorHandling:** Ensures the service correctly handles invalid input data and returns appropriate error messages and codes.

## Example

1. Process Receipt

   ````
   curl -X POST \
     http://localhost:8080/receipts/process \
     -H 'Content-Type: application/json' \
     -d '{
       "retailer": "Supermarket",
       "purchaseDate": "2024-03-19",
       "purchaseTime": "15:00",
       "items": [
         {
           "shortDescription": "Apples",
           "price": 3.50
         },
         {
           "shortDescription": "Bread",
           "price": 2.25
         }
       ],
       "total": 5.75
   }'
   ````

2. Calculate Points using ID

   ````
      curl -X GET \
      http://localhost:8080/receipts/{id}/points
   ````


To run these commands, execute them in your terminal. Ensure that your receipt processing service is running and replace ```http://localhost:8080``` with the actual base URL of your service.

Alternatively, you can import these curl commands into Postman:

Open Postman.
- Click on the "Import" button.
- Choose "Raw Text" and paste the curl commands.
- Click on "Continue" and then "Import."
- You can now execute these requests directly from Postman by clicking the "Send" button.
These curl examples demonstrate how to process a receipt and calculate points using the provided API endpoints.
