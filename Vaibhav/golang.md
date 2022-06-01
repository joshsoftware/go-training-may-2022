## Why Golang ?
---

* Go was originally designed by google in 2007.The go programming language is an open-source programming language and go language has focused on simplicity , reliability and efficiency.

* Go was originally designed for networking and infrastructure related programs.There are many high performance programming languages other than golang.And go was developed to replace these languages.


  ### Features
 * Concurrency :- Can handle millions of platform users.

 * Scalability :- Grows with the business to handle severals simultaneous task.

 * Compiled Language :- Fast performance as it converts directly into machine-level code.

 * Cross-Platform :- Golang perform well in across all platforms.

   ### Application 

 * Using go we can design server side applications.
 * Using go-lang many command line tools are developed.
 * GO can be used in AI and Data Science world.
 * Using go we can also develope a game.

--- 
These are some points i considered while choosing Golang as Programming language.

---

## GOROOT and GOPATH 

* GOPATH & GOROOT are the environment variable.
	
### what is environment variable ?
- Environmental variables are variables that are defined for the current shell and are inherited by any child shells or processes. 		-	Environmental variables are used to pass information into processes that are spawned from the shell.

    > ## GOROOT
    * $GOROOT has always been defined as a pointer to the root of your Go installation.
     * The Go binary distributions assume they will be installed in /usr/local/go

    > ## GOPATH
    * GOPATH, also called the workspace directory, is the directory where the Go code belongs.

    * It is implemented by and documented in the go/build package and is used to resolve import statements. 

---
## Day 2
  * Data type is an important concept in programming. Data type specifies the size and type of variable values.

   * Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

  * Go has three basic data types:

    > bool: represents a boolean value and is either true or false

    >Numeric: represents integer types, floating point values, and complex types

    >string: represents a string value
---
## Day 3

  > ### Function

     A function is a group of statements that together perform a task

     A function declaration tells the compiler about a function name, return type, and parameters. A function definition provides the actual body of the function.

> ### Loops

    * when you need to execute a block of code several number of times then we use loop.

    * A loop statement allows us to execute a statement or group of statements multiple times.

> ### Pointer

     * A pointer is a variable whose value is the address of another variable


    var var_name *var-type


## Day 4
    
  > ### Arrays

    * In Go, an array is a numbered sequence of elements of a specific length. 

  >Declaring Arrays :-

  * To declare an array in Go, a programmer specifies the type of the elements and the number of elements required by an array as follows.

        var variable_name [SIZE] variable_type

  > ### Slices

    * Slices are similar to arrays, but are more powerful and flexible.

    * Like arrays, slices are also used to store multiple values of the same type in a single variable.

    * However, unlike arrays, the length of a slice can grow and shrink as you see fit.

    * In Go, there are several ways to create a slice:

    * Using the []datatype{values} format

    > slice_name := []datatype{values}
    > myslice := []int

*  in Go, there are two functions that can be used to return the length and capacity of a slice:

    >len() function - returns the length of the slice (the number of elements in the slice)

    > cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)



