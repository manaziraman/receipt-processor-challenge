# Receipt Processor Challenge Guide

This README provides instructions for building, running, and testing the application using Docker. It also breaks down
my design decisions when developing this backend app.

# Running the Application
- `make all`: *Use this command to run the app*. It builds the Docker image and runs the application in a container in one step. It's a shortcut for make build followed by make run.
- `make build`: This command builds the Docker image for the application using the specified IMAGE_NAME. It uses the Dockerfile in the current directory to define how the image should be built.
- `make run`: This command first builds the Docker image (if not already built) and then runs the application inside a Docker container. The container is named according to the CONTAINER_NAME, and the application is accessible at http://localhost:8080.
- `make test`: This command runs the tests for the application using the local Go installation. It does not use Docker to run the tests.
- `make stop`: This command stops the running Docker container with the name specified in CONTAINER_NAME. If the container is not running, it will not produce an error.
- `make remove-container`: This command removes the stopped Docker container with the name specified in CONTAINER_NAME. If the container does not exist, it will not produce an error.
- `make clean`: This command stops and removes the running container and then removes the Docker image specified by IMAGE_NAME. It's a complete cleanup command.
- `make rebuild`: This command performs a full cleanup (using make clean), then rebuilds the Docker image and runs the application in a new container. It's useful when you want to start fresh with a new build and container.

# Design Decisions
- The application has been thoughtfully structured into different components such as handlers, models, receipts, and database to promote modularity, maintainability, and separation of concerns. By dividing the app into these specific parts, each component can be developed, tested, and maintained independently, leading to a more robust and scalable system.
- The inclusion of tests ensures that the core functionality is validated, providing confidence in the code's correctness and facilitating future enhancements or refactoring.
- The use of a Makefile streamlines the build and deployment process, making it easier for developers to execute common tasks without having to remember complex command-line instructions. This contributes to a more efficient development workflow.
- Regarding the decision not to include concurrency checks for the database, it was determined that they would only be necessary if there were a likelihood of collision in IDs. Since UUIDs are used in the system, the chances of such collisions are extremely unlikely, rendering the additional complexity of concurrency checks unnecessary. This decision reflects a balance between robustness and simplicity, taking into consideration the specific requirements and constraints of the application.