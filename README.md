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