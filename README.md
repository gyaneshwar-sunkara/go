# GO

- In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.

- Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).

- In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type).

- Any Go function can return multiple values.

- Go slice. A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types.

- When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the size of the array underlying the slice can be dynamically changed

- Changing the Hello function's parameter from a single name to a set of names would change the function's signature. If you had already published the example.com/greetings module and users had already written code calling Hello, that change would break their programs.

  In this situation, a better choice is to write a new function with a different name. The new function will take multiple parameters. That preserves the old function for backward compatibility.

- In Go, you initialize a map with the following syntax: make(map[key-type]value-type).

- In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value.

  You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.

- Ending a file's name with \_test.go tells the go test command that this file contains test functions.

- Test functions take a pointer to the testing package's testing.T type as a parameter.

  You use this parameter's methods for reporting and logging from your test.

-

### Commands

`go mod init example.com/hello`

`go run .`

`go help`

`go mod tidy`

`go mod edit -replace example.com/greeting=../greeting`

`go build`

`go list -f '{{.Target}}'`

`go install`
