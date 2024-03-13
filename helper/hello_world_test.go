package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Thomas", func (t *testing.T) {
		result := HelloWorld("Thomas")
		require.Equal(t, "Hello Thomas", result, "Result must be 'Hello Thomas'")
	})

	t.Run("Ardiansah", func (t *testing.T) {
		result := HelloWorld("Ardiansah")
		require.Equal(t, "Hello Ardiansah", result, "Result must be 'Hello Ardiansah'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Thomas")
	require.Equal(t, "Hello Thomas", result, "Result must be 'Hello Thomas'")
}

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