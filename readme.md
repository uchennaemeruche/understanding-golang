# Golang Learning Files - Introduction

This repo houses some Golang introductory files, sample codes and implementations. I will be updating it as I keep getting a hang of the language.

# Author: Uchenna Emeruche

- [Section 1: Golang Introduction and Setup(Using Docker)](#golang-introduction-and-setup)
- [Section 2: Declaring variables and Assignments in Golang]
- [Section 3: Data types in Golang - Strings, Numbers, Arrays, Slices, Maps, Interfaces etc]
- [Section 4: Golang Control Structures: Looping and Iteration in Golang]
- [Section 5: Golang Control Structures: Booleans and Conditionals in Golang (If/else, switch statements)]
- [Section 6: Functions in Golang]
- [Section 7: Reciever Functions in Golang]
- [Section 8: Structs and Custom Types in Golang]
- [Section 9: Structs and Custom Types in Golang]
- [Section 9: Interfaces in Go (Polymorphism, )]

# Golang Introduction and Setup

Golang(Go) is a strongly-typed, general-purpose, compiled programming language supported by Google which has built-in concurrency and a robust standard library. Go is an open source language built to be high-performant and efficient.

To download and install Go on your computer, head over to [Download and Install Go](https://go.dev/doc/install). Click the button to download Go for your Operating System and follow the installation instruction on the page. After installation, open up a terminal and type `go run ` to test your go installation.

> Whereas Go can be installed and Run on your local computer, in this Learning tour, I've chosen to run Go using docker. You don't have to follow this route but I'll recommend it if you are big on the DevOps practice. By containerizing your applications, you can reap the benefit of **Portability**: deploy to any other system where Docker is running and be sure that your app will perform exactly as it did when you tested it locally.
>
> > Ensure you have Docker installed and running on your machine to follow through with this method. You can read up on that here: [Docker installation and confifuration](https://docs.docker.com/engine/install/).

Depending on the folder structure you want to maintain, I have created a root folder called **golang** where I'll be creating all my Go applications and learning repos. For this section, I will be working in a folder called **understanding-go** inside the golang directory. Go ahead and clone the repo.
The **understanding-go** directory contains a **dockerfile** for building our docker image. There is nothing fancy in the file except that I'm pulling down _golang:1.17_, specifying the build stage as _dev_ and setting up my _WORKDIR_ as /dev

First thing is to change directory to the **understanding-go** folder `cd understanding-go` and build the docker image from the dockerfile as follows:

```
docker build --target dev . -t go-image
```

Similarly, we can run the container in interactive mode as follows:

```
    docker run -it -v ${PWD}:/dev go-image sh
```

The above command will take you into the container where you can now start running your go applications.

> TO confirm that Golang has been installed on the container run `go version`
