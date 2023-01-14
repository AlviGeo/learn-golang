package helper

import (
	"fmt"
	"runtime"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name: "Alvi",
			request: "Alvi",
		},
		{
			name: "Geovanny",
			request: "Geovanny",
		},
		{
			name: "AlviGeovanny",
			request: "AlviGeovanny",
		},
	}

	for _,benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Alvi", func(b *testing.B) {
		for i := 0; i<b.N; i++ {
			HelloWorld("Alvi")
		}
	})
	b.Run("Geovanny", func (b *testing.B)  {
		for i := 0; i<b.N; i++ {
			HelloWorld("Geovanny")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Alvi")
	}
}

func BenchmarkHelloWorldGeovanny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Alvi")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	} {
		{
			name: "Alvi",
			request: "Alvi",
			expected: "Hello Alvi",
		},
		{
			name: "Geovanny",
			request: "Geovanny",
			expected: "Hello Geovanny",
		},
		{
			name: "Zoey",
			request: "Zoey",
			expected: "Hello Zoey",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Alvi", func(t *testing.T) {
		result := HelloWorld("Alvi")
		require.Equal(t, "Hello Alvi", result, "Result must be 'Hello Alvi'")
	})
	t.Run("Geovanny", func(t *testing.T) {
		result := HelloWorld("Geovanny")
		require.Equal(t, "Hello Geovanny", result, "Result must be 'Hello Geovanny'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

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