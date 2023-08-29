# mirror
Golang helper for the reflect package.

```sh
go get github.com/mark-veres/mirror
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/mark-veres/mirror"
)

type User struct {
	Name string
	Age  int
}

func (u *User) CanDrive() bool {
	return u.Age >= 18
}

func main() {
    s := User{
        Name: "bob",
        Age: 18,
    }

    // returns an array with
    // the name of all fields
    fields := mirror.Fields(s)

    // returns an array with
    // the name of all methods
    methods := mirror.Methods(&s)

    // changes the value for "Name"
    // field to "mark". returns an error
    err := mirror.SetField(&s, "Name", "mark")

    // returns a reflect.Value object and an error
    value, err := mirror.GetField(&s, "Name")
    fmt.Println(value.String())
}
```