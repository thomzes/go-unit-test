package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Thomas", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Thomas")
		}
	})

	b.Run("Ardiansah", func (b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ardiansah")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Thomas")
	}
}

func BenchmarkHelloWorldArdiansah(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ardiansah")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "Thomas",
			request: "Thomas",
			expected: "Hello Thomas",
		},
		{
			name: "Ardiansah",
			request: "Ardiansah",
			expected: "Hello Ardiansah",
		},
		{
			name: "Dany",
			request: "Dany",
			expected: "Hello Dany",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func (t *testing.T)  {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

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