package main

import (
	"fmt" // this is part of Go's standard pacakge
	_ "time"
)

func main() {
	//println("Hello World")
	fmt.Println("Hello World")
}

// package must be main to write main function
// There are neither arguments nor a return for main
// ctrl + shift + p
// to run go run hello.go
// to build go build hello.go
// to build and generate a binary/exe with a desired name then go build -o app hello.go
// to install binaries to go bin go install hello.go
// cross compilation and build for different machines
// GOOS=windows GOARCH=amd64 go build -o app.exe hello.go

// To know all supported platforms aka targeted platforms in go
// go tool dist list
// android/arm64 --> GOOS is android. GOARCH is arm64
// darwin/amd64  --> GOOS is darwin.  GOARCH is amd64
// windows/amd64

// there three types of packages in go
// 1- Go standard packages (std library)
// 2- User defined packages (Packages those are written by you or your teammates)
// 3- Third party packages ususlly pacakges those are downloaded from github or similar systems.

// Imported and not used is not a warning in Go. It is an error
// To import and to use it later you can declare it or named it using a blank identifier _
