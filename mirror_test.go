package mirror_test

import (
	"fmt"
	"testing"

	"github.com/mark-veres/mirror"
)

type User struct {
	Name string
	Age  int
}

func (u *User) CanDrive() bool {
	return u.Age >= 18
}

func TestFields(t *testing.T) {
	s := User{
		Name: "bob",
		Age:  18,
	}

	fields := mirror.Fields(s)
	correct := fields[0] == "Name"
	correct = fields[1] == "Age" && correct

	if !correct {
		t.Fail()
	}
}

func TestMethods(t *testing.T) {
	s := User{
		Name: "bob",
		Age:  18,
	}

	methods := mirror.Methods(&s)

	if methods[0] != "CanDrive" {
		t.Fail()
	}
}

func TestSetField(t *testing.T) {
	s := User{
		Name: "bob",
		Age:  18,
	}

	err := mirror.SetField(&s, "Name", "mark")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if s.Name != "mark" {
		t.Fail()
	}
}

func TestGetField(t *testing.T) {
	s := User{
		Name: "bob",
		Age:  18,
	}

	value, err := mirror.GetField(&s, "Name")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if value != "bob" {
		t.Fail()
	}
}
