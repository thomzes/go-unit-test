package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Thomas")

	assert.Equal(t, "Hello Thomas", result, "Result must be 'Hello Thomas'")

	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Thomas")

	require.Equal(t, "Hello Thomas", result, "Result must be 'Hello Thomas'")

	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Thomas")

	if result != "Hello Thomas" {
		// error
		t.Error("result must be Hello Thomas")
	}

	fmt.Println("TestHelloWorld Done!")
}

func TestHelloWorldArdiansah(t *testing.T) {
	result := HelloWorld("Ardiansah")

	if result != "Hello Ardiansah" {
		// error
		t.Fatal("result must be Hello Ardiansah")
	}

	fmt.Println("TestHelloWorldArdiansah Done!")
}