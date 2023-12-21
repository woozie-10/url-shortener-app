# URL Shortener App
This web application serves as a URL Shortener, built with Go and Gin. Users can easily shorten URLs through a user-friendly interface.
## Technologies

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [API](https://cleanuri.com/docs)
- [Docker](https://www.docker.com/)
- [Viper](https://github.com/spf13/viper)

## Project Structure

The project is organized into the following main components:

- `cmd`: Contains the main application entry point.
- `router`: Manages Gin router settings for efficient HTTP routing and request handling in the web application.
- `api`: Contains the implementation of client-server interaction.
- `entities`: Defines the data model for the `Response` entities.
- `handlers`: Implements the HTTP request handlers.
- `config`: Contains configuration settings and management for the web application.
- `tests`: Responsible for writing and running test cases to ensure the correctness and reliability of application.
- `frontend`: Contains the client side of the application

## Launching the Application

```shell
# Clone the repository:
git clone https://github.com/woozie-10/url-shortener-app.git

# Change to the project directory:
cd url-shortener-app

# Run the application:
make build && make run

```
The server should start, and you can access the app at http://localhost:2311.

# Running Tests in a project

To run tests in a project, follow these steps:

1. **Navigate to Project Root**: First, open your terminal and navigate to the root directory of project:

   ```shell
   cd url-shortener-app
   ```
3. **Run Tests**: Run the tests using the command. This command will automatically discover and execute all the test functions in project:

   ```shell
   make test
   ```
