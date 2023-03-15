package employee

import (
	"fmt"
	"testing"
)

var Debug bool = false

func TestNewFullName(t *testing.T) {
	firstName := "John"
	lastName := "Doe"
	if Debug {
		t.Skip("skip test")
	}

	v, err:= NewFullName(firstName, lastName)
	if err!= nil {
        t.Fatal(err)
    }

	fmt.Println(v)

	// if v == {FirstName: "John", LastName: "Doe"} {
	// 	t.Error("fullName is not valid")
	// }
}

// func TestNewFullName(t *testing.T) {
// 	fullName, err := NewFullName("john", "Doe")
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	if fullName.FirstName() != "john" {
// 		t.Errorf("expected firstName to be john, but got %s", fullName.FirstName())
// 	}

// 	if fullName.LastName() != "Doe" {
// 		t.Errorf("expected lastName to be Doe, but got %s", fullName.LastName())
// 	}

// 	_, err = NewFullName("", "Doe")
// 	if err == nil {
// 		t.Errorf("expected error, but got nil")
// 	}

// 	_, err = NewFullName("john", "")
// 	if err == nil {
// 		t.Errorf("expected error, but got nil")
// 	}
// }
