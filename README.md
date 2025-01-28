# task-manager-golang

RESTful API service built in Go that provides task management capabilities.

Installation Guide
To install Task Manager Golang, follow these steps:

Step 1: Install Go
Task Manager Golang is built using Go, so you need to have Go installed on your system. You can download the latest version of Go from the official Go website (https://golang.org/dl/). Follow the installation instructions provided on the website to install Go on your system.

Step 2: Install Go Modules
Task Manager Golang uses Go Modules to manage dependencies. To use Go Modules, you need to have Go 1.11 or later version installed on your system. If you haven't installed Go 1.11 or later, install Go using the instructions provided in step 1 and then follow the instructions below.

Step 3: Clone the Repository
Clone the Task Manager Golang repository from GitHub using the following command:

```shell
git clone https://github.com/task-manager-golang/
task-manager-golang.git
cd task-manager-golang
```

Step 4: Install all the dependencies

```shell
go mod tidy
```

Step 5: Set up env variables. Create a `.env` file in the root of the project:

```env
PORT=3001
DB_URL="host=localhost user=dbuser password=dbpass dbname=postgres port=5432 sslmode=disable"
```

Step 6: Initialise the tables

```
CREATE TABLE tasks (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(255) NOT NULL
);
```

Step 7: Build and Install Task Manager Golang
Navigate to the cloned repository and run the following command to build and install Task Manager Golang:

```shell
go build -o task-manager-golang
```

This command will build the Task Manager Golang application and create a binary file named task-manager in the current directory.

Step 8: Run Task Manager Golang
To run Task Manager Golang, simply run the binary file using the following command:

```shell
./task-manager-golang
```

Task Manager Golang will start and you can use it to manage your tasks.

Note:

Make sure you have Go installed on your system before proceeding with the installation.
Ensure that you have Go Modules installed on your system if you want to use it to manage dependencies.
Clone the Task Manager Golang repository from GitHub using the provided command.
Build and install Task Manager Golang using the provided command.
Run Task Manager Golang using the provided command to start using it.
That's it! You have successfully installed and configured Task Manager Golang on your system.
