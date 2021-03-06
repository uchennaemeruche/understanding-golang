# Golang Learning Files - Introduction

This repo houses some Golang introductory files, sample codes and implementations. I will be updating it as I keep getting a hang of the language.

# Author: Uchenna Emeruche

- [Section 1: Golang Introduction and Setup(Using Docker)](#golang-introduction-and-setup)
- [Section 2: Declaring variables and Assignments in Golang]
- [Section 3: Types in Golang - Strings, Numbers, Arrays, Slices, Maps, Interfaces etc]
- [Section 4: Golang Control Structures: Looping and Iteration in Golang]
- [Section 5: Golang Control Structures: Booleans and Conditionals in Golang (If/else, switch statements)]
- [Section 6: Functions in Golang]
- [Section 7: Receiver Functions in Golang]
- [Section 8: Structs and Custom Types in Golang]
- [Section 9: Interfaces in Go ](#interfaces-in-go)

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

# Interfaces in Go

Interface is one of Go's implementation of class inheritance. The other being **embedding**. Embedding is a limited form of inheritance which allows types to share data and code. Embedding can be viewed as an automated form of composition

Interfaces are types made up of named collection of method signatures. In Go, interfaces provide a way to specify the behaviour/methods of an object: _If something can do this, then it fits in here_. They define contracts that certain objects or types should adhere to. Interfaces define what methods a type should have and the type decides how to implement these methods.

> Golang ships with a nominal type system, however, Interfaces provide a limited form of structural typing(In a structural type sytem, compatibility and equivalence are determined by the actual structure or definitions of the type).
>
> > This is a core concept in Go???s type system; instead of designing our abstractions in terms of what kind of data our types can hold, we design our abstractions in terms of what actions our types can execute.

When a type provides definitions for all the methods in an interface, it is said to implement the interface.

> To implement an interface in Go, we only need to implement all the methods in the interface. Note that there is no **implement** keyword in Go; whether or not a type satisfies an interface is determined automatically.

For example, A Car interface can define certain methods such as Drive(), Park() and Reverse(). Now any type that can provide definitions for these methods satisfies the Car interface and that **type** is said to implement the Car interface:

```
type Car interface {
	Drive() string
	Park() string
	Reverse() bool
}

type Toyota struct {
	color string
	model string
	speed int
}

func (t Toyota) Drive() string {
	str := fmt.Sprintf("Your %v Toyota car is running on %vkm/hr", t.model, t.speed)
	return str
}

func (t Toyota) Park() string {
	elapsedTime := 10.0
	distanceTravelled := float64(t.speed) * elapsedTime
	str := fmt.Sprintf("%v %v Toyota covered %vkm", t.color, t.model, distanceTravelled)
	return str
}

func (t Toyota) Reverse() bool {
	return true
}

func main() {
	cars := []Toyota{
		{color: "red", model: "2007", speed: 70},
		{color: "blue", model: "2010", speed: 120},
	}

	for _, car := range cars {
		fmt.Println(car.Drive())
		fmt.Println(car.Park())
		fmt.Println(car.Reverse())
	}

}

```

In the above snippet, the Toyota type implements the Car interface by providing definition for all the method set in the interface.
i.e By adding the methods Drive(), Park() and Reverse() to the receiver type _Toyota_, the Toyota type is said to implement the Car interface.

## The interface{} type (empty interface)

The interface{} type, (empty interface) is the interface that has no methods.

> Since there is no **implements** keyword in Go, all types implement at least zero methods, and since satisfying an interface is done automatically, all types satify the empty interface.
>
> > That means given a function that takes an interface{} as a parameter, you can pass any value to that function when calling it.

```
func Foo(param interface{}){

}

```

Looking at the function above, one would say that the parameter _param_ is of any type. I mean the empty interface _interface{}_ is supposed to represent a generic, right!.

> > OOps!! Not really, I mean not so fast. The value passed to the function must in fact be of type interface{}

This is very important to note because at runtime, all values have exactly one type. So under the hood, the Go runtime will try to perform a type conversion and convert the value to an interface{} type.

This article [Go Data Structures: Interfaces](https://research.swtch.com/interfaces) explains the conversion concept in details.
But the key takeaway is that interface values are represented as a two-word pair. The first being a pointer to the metadata about the type stored in the interface. And the second being a pointer to the actual data held by the value.
