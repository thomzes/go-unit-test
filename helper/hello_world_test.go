package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Thomas")
	if result != "Hello Thomas" {
		// error
		panic("Result is not Hello Thomas")
	}

}

func TestHelloWorldArdiansah(t *testing.T) {
	result := HelloWorld("Ardiansah")
	if result != "Hello Ardiansah" {
		// error
		panic("Result is not Hello Ardiansah")
	}

}