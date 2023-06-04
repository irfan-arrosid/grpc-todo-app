# gRPC Todo App

Simple Todo App using gRPC for server-client communication. This project's tech is Golang, PostgreSQL, and GORM.

## Prerequisites:
- Make sure you have Go and PostgreSQL installed on your computer.
- Create a PostgreSQL database and adjust the database connection configuration in the `.env` file according to your database settings.

## Install Dependencies:
- Open a terminal and navigate to the project directory.
- Run the command `go mod tidy` to install all the required dependencies.

## Run the Server:
- Navigate to the server directory using command `cd server`.
- Execute the command `go run server/main.go` to run the gRPC server.
- The server will start running and will be ready to accept requests from the client.

## Run the Client:
- Open a new terminal and navigate to the client directory using command `cd client`.
- Execute the command `go run client/main.go` to run the gRPC client.
- The client will connect to the server and display a menu of options for todo operations.

This project can also testing using Postman by importing the .proto file.
