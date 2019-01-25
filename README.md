# CARS

This project is just a simple example of REST API created in Golang.

It will show you list of the cars, stores in CSV file.

## How to use it:
1. Clone this repo
2. Run it using `go run` tool:
```sh
 $ go run main.go
```
3. Visit `http://localhost:8080/cars` to check the list of the cars

## Prerequisites:
You just have to install [Golang](https://golang.org/doc/install) and [mux](https://github.com/gorilla/mux) to route requests:
```sh
$ go get github.com/gorilla/mux
```