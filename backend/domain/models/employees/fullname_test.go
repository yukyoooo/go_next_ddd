package employee

import (
	"testing"
)

func TestNewFullName(t *testing.T) {
	fullName, err := NewFullName("john", "Doe")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if fullName.FirstName() != "john" {
		t.Errorf("expected firstName to be john, but got %s", fullName.FirstName())
	}

	if fullName.LastName() != "Doe" {
		t.Errorf("expected lastName to be Doe, but got %s", fullName.LastName())
	}

	_, err = NewFullName("", "Doe")
	if err == nil {
		t.Errorf("expected error, but got nil")
	}

	_, err = NewFullName("john", "")
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
