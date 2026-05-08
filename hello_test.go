package hellogo_test

import (
	"fmt"
	"github.com/PeterChen607/hellogo"
	"testing"
)

func TestHello(t *testing.T) {
	data := "jack"
	expected := fmt.Sprintf("hello %s!", data)
	result := hellogo.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s", expected, result)
	}

}
