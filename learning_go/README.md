# Learning How to Code: Google's Go (golang) Programming Language

## Course Resources

[ArdenLabs](https://www.ardanlabs.com/)

[Course Google Docs](https://goo.gl/PHKgO7)

[Presentations](https://drive.google.com/drive/folders/0B22KXlqHz6ZNfjNXTzk1U3JHUkJ6VjJ3dnJKNzVtNjRUM3Q2WFNqWGI2Q3RadERqUlVrOEU)

[Todd McLeod's youtube channel](https://www.youtube.com/user/toddmcleod)

[Go Lang Forum](https://forum.golangbridge.org/)

[Dave Cheney's Blog](https://dave.cheney.net/)  

## Installing Go

### Environment Variables

Type go env to show Go specific environment variables.

Two required environment variables:  

    GOROOT is set to where you installed GO (e.g., Windows GOROOT=c:\GO)

    GOPATH is set to your workspace

### Go Editors

Course instructor likes WebStorm, but he talks about Atom

## The Go command and Documentation

### Finding Help

At command line type "go help"
golang.org

## How Computers Work

* Computers run on electricity
* Circuit/Transistor/Switch - something that is on or off.
* Coding schemes can be created based on a one or more circuits and their respective states

Todd's Coding Scheme

off : off | Go away  
off : on  | Come in  
on  : off | Bring Pizza  
on  : on  | Bring beer  

If we have three lights we have 2^n or 2^3 choices.

TODO: Review binary, hex, and character representation.

## Numeral Systems

A writing system for expressing numbers.

* Binary
* Hexadecimal
* Decimal

## Section 5

### Packages

Folder name needs to be same as package name.
No main file in package.

All files in package folder must have the same package name.

Package scope: files in same package can call functions from other files in that same package.

Lower case funcs are not exported outside the package, they are visible inside the package.

Can't have more than one func main() {} in main package.

Useful go commands:

* go run
* go build
* go install
* go clean

Declaring varialbes:

* var myName = "Bill"
* var myName string = "Bill"
* myName := "John"

### Go Commands

go run [filename] - Runs the file.
go build - Builds executable if folder has executable.
go fmt - Format go code in current folder.
go clean - Cleans current folder (e.g., deletes executable)
go install - If in main, packages and installs the main executable in workspace bin folder.

### Variables

Many ways, but two methods are preferred:

* shorthand
  * can only be used inside func  

  * handles declaration, 
    assignment, initialization

  * Examples:
        a := 10
        b := "golang"
        c := 4.17
        d := true
* var - zero value

### Scope

* Universe
* Package - outside of any containng block.  Only accessible to whole pacakge.
* File
* Block (curly braces)

Order does not matter in package level.  Order matters in blocks.
Importing is at the file level.  

```go
// Package level

package main

import "fmt"

var x int = 42 // Package accessible var

func main() {
    y : = 7 // Block level scope.  Only availabel in func main

    fmt.Println(x)
}
```

```go
// Example of shadowing

package main

import "fmt"

func max(x int) int {
    return 42 + x
}

func main() {
    max := max(7)
}
```

Declare vars as near to use as possible.

### Closure

```go
package main

import "fmt"

var x int = 42

func main() {
    // Outer scope

    fmt.Println(x)
    foo()

    // Closure    
    fmt.Println(x)
    {
        // Inner scope 
        fmt.Println(x)
        y := "The credit belongs with theone who is in the ring."
        fmt.Println(y)
    }

    // An anonymous function - functions without name.
    // A function expression
    increment := func() int {
        x++
        return x
    }

    increment := wrapper()
}

func wrapper() func() int {
    x := 0  // Always remembered by outer scope
    return func() int {
        x++
        return x
    }

}

func foo() {
    fmt.Println(x)
}
```

### Constants I

Best practices, constants are lower case.

Multiple assignment example:  

```go
const (
   pi = 3.14
   language = "go"
)
```

*iota*  - An extremely small amount.  

```go
const (
    A = iota // 0
    B // 1
    C // 2
)

const (
    D = iota // 0
    E // 1
    F // 2
)
```

### Constants II

[Rob Pike's blog post](https://blog.golang.org/constants).  

* Go has NO mixing of numeric types.  
* Statically typed.  
* Types are checked at compile time.  

Constants can be _typed_ and _untyped_.  

   const hello = "Hello" // untyped
   const hello string = "hello" // typed

* Untyped - const value that does not yet have a fixed type, a "kind", not yet forced to obey strict
rules that prevent combining differently typed values.

x := 42 // 42 is untyped (aka "kind").  

### Memory Addresses

Post Office boxes analogy.  
Memory slots/addresses.  

Show var's memory address by using ampersand.
fmt.Println("Address: ", &a)  // Displays memory hex value

## Control Flow

* Code read in sequence, top to bottom  
* Loops/Iteration  
* Scope  
* Conditionals (e.g., if-then-else, switch)

Switch does NOT have the typical default fall-through.  

* For-loop

### Documentation and Terminology

Rune - a single character (any language).  Also, an alias for int32.  An integer value identifying a unique unicode code point.  

## Switch

Like other languages, but does not have typical fall-through.  Have to explicitly use fallThrough if want fall-through behavior.  Break is default behavior for switch statement.  

```go
switch x.(type) {
    case int:
        // do something
    case string:
        // do something else
    case Person:
        // do person thing
    ...
}
```

## If-then-else

If-then-else and switchs accept or work with [initialization statements](https://golang.org/doc/effective_go.html#if).  

## Functions

func main() {} - is the entry point to a program.  Can only have one func main in entire program and func main has to be in package "main".  

[Function Language Spec](https://golang.org/ref/spec#Function_types)  

### Variadic Functions

The final parameter in a func signature can have zero or more args if that param is declared with the "...", variatic, notation.  

## Callbacks

Functions are types in the golang spec.  That means that a function can be passed as args to functions.  Or, a function can be returned by a function.  

```go
// Example

package main

import "fmt"

funct visit(numbers []int, callback func(int)) {
    for _, n = range numbers {
        callback(n)
    }
}

func main() {
    visit([]int{1, 2,3, 4}, func(n int){
        fmt.Println(n)
    })
}
```

### Defer

Defer is a keyword.  [Defer statements](https://golang.org/ref/spec#Defer_statements) invoke a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding fuction executed a return, reached the end of its function body or because the correspoinding goroutine is panicking.

Typically use defer to handle closing files.  

### Pass By Value

Everything in Go is passed by value.  

### Anonymous Self Executing Function

```
package main

import "fmt"

func main() {
    func() {
        fmt.Println("I'm here.")
    }
}
```

## Section 40 - Slices

Slices are lists.  

```go
myslice := []int{1,2,3,5} // A slice of int's with values.
myslice[2:4]  // Slicing a slice.  Give everthing from position 2 up to, but not including position 4
myslice[2]    // Get item at position 2
"myString"[2] // Item at position 2 in "myString", which is S, aka Index access.
```

Slices have length, the number of items in the slice.
Slices have capacity, the number of items the slice can contain.

Creating a slice: make([]T, length, capacity)

```go
mySlice := make([]int, 50, 100)
mySlice := new([100]int)[0:50]
```

### More on Slices

## Secton 41 - Data Structures - Map

Key/value based data structure.  Sometimes known as dictionaries in other programming languages.  

* slices - for lists
* maps - for key-value pairs, unordered, built on top of a hash table
* structs - for composite data

* "comma ok" idiom - use to see if value is assigned or empty.

```go
var seconds int
var ok bool
seconds, ok = timeZone[tz]

for offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }

    log.Println("Uknown time zone: ", tz)
    return 0
}
```

## Hashtables

[One of the most useful data structures](https://en.wikipedia.org/wiki/Hash_table).  Enable very fast searching.  

## Creating First Project with WebStorm

* Package names should be short, concise and provocative.  
* Folder name should be same as package name.  
* Can see go commands just by typing "go" at the command prompt.
  * go fmt - will format your go code.
  * go run - will run the go file.
  * go build - will build the program executable to the folder name holding the code, puts executable in workspace/bin
    * $GOPATH/bin
  * go clean - cleans the current workspace
  * go env - show go's environment variables/settings
  
