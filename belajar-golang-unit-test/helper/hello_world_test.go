package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on Mac OS")
	}

	result := HelloWorld("Alvi")
	require.Equal(t, "Hello Alvi", result, "Result must be 'Hello Alvi'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Alvi")
	require.Equal(t, "Hello Alvi", result, "Result must be 'Hello Alvi'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Alvi")
	assert.Equal(t, "Hello Alvi", result, "Result must be 'Hello Alvi'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAlvi(t *testing.T) {
	result := HelloWorld("Alvi")
	if result != "Hello Alvi" {
		// panic("Result is not Hello Alvi")
		t.Error("Result must be 'Hello Alvi'")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldGeovanny(t *testing.T) {
	result := HelloWorld("Geovanny")
	if result != "Hello Geovanny" {
		// t.FailNow()
		t.Fatal("Result must be 'Hello Geovanny'")
	}

	fmt.Println("TestHelloWorldGeovanny Done")
}