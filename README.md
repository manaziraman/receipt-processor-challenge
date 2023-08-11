# Receipt Processor Challenge - Backend Software Engineer

A backend REST API web service that accepts HTTP requests and returns responses based on conditions outlined in the section below. The application calculates points based on the receipt details and provides endpoints to process receipts and retrieve points.

Visit https://github.com/fetch-rewards/receipt-processor-challenge for more detail on examples and the project specifications.

### Premise
The application processes receipts and calculates points based on specific criteria:

- Points for alphanumeric characters in the retailer name.
- Bonus points for round dollar amounts, multiples of 0.25, and specific purchase times.
- Additional points based on item descriptions and prices.

### Dependencies/Tools
- Go - Open-source programming language that makes it easy to build simple, reliable, and efficient software.
- Mux - A powerful HTTP router and URL matcher for building Go web servers.
- Docker - A platform used to develop, ship, and run applications inside containers.
- Makefile - A simple way to manage build details.

## Getting Started
You must have Go version 1.17 or higher and Docker installed.

1. Verify Go version
```
go version
```
2. Clone Repository Locally
```
git clone https://github.com/manaziraman/receipt-processor-challenge.git
```
3. Go to the project's root directory
```
cd /my/path/to/receipt-processor-challenge
```
4. Build and run the application using Makefile
```
make all
```
5. Your terminal should read:
```
Server started at http://localhost:8080
```
6. Verify the app is running by visiting http://localhost:8080

## Making API Calls
You can use tools like Postman or Curl to make calls to the API.

### Process Receipt
- **POST** Route "/receipts/process"
- REQUEST BODY FORMAT:
```
{
  "retailer": <str>,
  "purchaseDate": <str>,
  "purchaseTime": <str>,
  "items": [<Item>],
  "total": <str>
}
```
### Get Points for Receipt
- **GET** Route "/receipts/{id}/points"
- Replace `{id}` with the UUID of the receipt.

## Running Tests
Run tests from the project's main directory:
```
make test
```

Tests check that the app should:

- Process a receipt and calculate points correctly.
- Handle various edge cases and error conditions.
- Retrieve points for a specific receipt.

## Design Decisions
The application is divided into handlers, models, receipts, and database to maintain a clean architecture. Handlers manage the HTTP request handling, models define the data structures, receipts contain the logic for processing, and the database handles data storage.

- Handlers: Separation of concerns for handling HTTP requests.
- Models: Definition of Receipt and Item structures.
- Receipts: Logic for processing receipts and calculating points.
- Database: In-memory storage of processed receipts.

Tests are included to ensure that the application behaves as expected, covering various edge cases and error conditions.

A Makefile is used to simplify the build and run process, encapsulating the necessary commands into simple make targets.

Concurrency checks for the database are omitted as UUIDs are used for receipt identification, making collisions extremely unlikely.

### Cleaning Up
To stop and remove the running container and image, execute:

```make clean```

This will stop the container, remove it, and remove the Docker image.

## Running the Application (Full Breakdown)
- `make all`: **Use this command to run the app**. It builds the Docker image and runs the application in a container in one step. It's a shortcut for make build followed by make run.
- `make build`: This command builds the Docker image for the application using the specified IMAGE_NAME. It uses the Dockerfile in the current directory to define how the image should be built.
- `make run`: This command first builds the Docker image (if not already built) and then runs the application inside a Docker container. The container is named according to the CONTAINER_NAME, and the application is accessible at http://localhost:8080.
- `make test`: This command runs the tests for the application using the local Go installation. It does not use Docker to run the tests.
- `make stop`: This command stops the running Docker container with the name specified in CONTAINER_NAME. If the container is not running, it will not produce an error.
- `make remove-container`: This command removes the stopped Docker container with the name specified in CONTAINER_NAME. If the container does not exist, it will not produce an error.
- `make clean`: This command stops and removes the running container and then removes the Docker image specified by IMAGE_NAME. It's a complete cleanup command.
- `make rebuild`: This command performs a full cleanup (using make clean), then rebuilds the Docker image and runs the application in a new container. It's useful when you want to start fresh with a new build and container.

## Design Decisions (Full)
- The application has been thoughtfully structured into different components such as handlers, models, receipts, and database to promote modularity, maintainability, and separation of concerns. By dividing the app into these specific parts, each component can be developed, tested, and maintained independently, leading to a more robust and scalable system.
- The inclusion of tests ensures that the core functionality is validated, providing confidence in the code's correctness and facilitating future enhancements or refactoring.
- The use of a Makefile streamlines the build and deployment process, making it easier for developers to execute common tasks without having to remember complex command-line instructions. This contributes to a more efficient development workflow.
- Regarding the decision not to include concurrency checks for the database, it was determined that they would only be necessary if there were a likelihood of collision in IDs. Since UUIDs are used in the system, the chances of such collisions are extremely unlikely, rendering the additional complexity of concurrency checks unnecessary. This decision reflects a balance between robustness and simplicity, taking into consideration the specific requirements and constraints of the application.