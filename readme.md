# BOOKSTORE

## ARCHITECTURE INFORMATIONS OF THE PROJECT

* Root directory /

### Pre-requisites

* Have goolang installed and configured on your machine.
* Have an IDE with support to Go programming language.

``` 
    git clone https://github.com/matcampos/bookstore.git
```

### Environments

before you run the project you need to configure the environment variable files.

Are them:

* .env - The .env.sample file shows all fields which this file need.
* docker.env - The same content of .env-sample file but with your docker container configurations, and DATABASE_HOST env value must be "database".

### Instalation

Execute the following command to install all dependencies of the project.

``` 
    go get -d -v ./...
```

To execute the project on your local machine run the following command on root directory of the project:

``` 
    go run main.go
```

### Build

To build the project execute the following command

``` 
    go install
```

- to run this build go on terminal in the root directory of the project and enter: `bookstore`

Another option to build and run the it is the following command:

``` 
    go build main.go && ./main.go
```

### Build and run with docker

To build the and run the project with docker you must configure the docker.env file, then you run the following command on root directory of the project.

``` 
    docker-compose up -d
```

## Built with

* [go1.15 darwin/amd64](https://golang.org/dl/) - Golang framework and go language download link.
* [errors](https://github.com/go-errors/errors) - Library to get the stackTrace of any error on go.
* [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Library to use MongoDB database on golang.
* [godotenv](https://github.com/joho/godotenv) - Library to use .env files and set the environment variable on golang.
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Library to validate structs properties.
* [gorilla/mux](https://github.com/gorilla/mux) - Library to create api endpoints on golang.

## Authors

* **Matheus Campos** - *Full-Stack Developer* - [Github](https://github.com/matcampos)
