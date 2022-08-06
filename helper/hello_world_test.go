package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Membuat Benchmark untuk mengukur kecepatan kode
func BenchmarkHelloWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld(" masendhy")
	}
}

//SubBenchmark
func BenchmarkSub(b *testing.B) {
	b.Run("sendhy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("sendhy ")
		}
	})

	b.Run("boedhi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("boedhi")
		}
	})
}

//Tabel Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "dekaqsa",
			request: "dek aqsa",
		},
		{
			name:    "razkadikr",
			request: "razka dzikr",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestHelloWorldMasendhy(t *testing.T) {
	result := HelloWorld(" masendhy")
	if result != " Hello masendhy" {
		//error
		//panic("Result isn't Hello masendhy") menggunakan kata kunci panic dalam test tidak dianjurkan.
		t.Fail() // meskipun test gagal, kode program di bawahnya tetap dijalankan, namun report terakhir tetap gagal
	}

	fmt.Println("TestHelloWorldMasendhy done")

}

func TestHelloWorldAbik(t *testing.T) {
	result := HelloWorld(" Abik")
	if result != "Hello Abik" {
		t.FailNow() // jika gagal kode programnya di hentikan
	}
	fmt.Println("TestHelloWorldAbik done")
}

// selain Fail dan FailNow ada juga Error dan Fatal, Error sama dengan FailNow tetapi bisa di beri argument didalam () nya, Fatal seperti Fail tetapi bisa diberikan argument

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld(" Abik")
	if result != "Hello Abik" {
		t.Error("Result must be Hello Abik") // jika gagal kode programnya di hentikan
	}
	fmt.Println("TestHelloWorldError done")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld(" Abik")
	if result != "Hello Abik" {
		t.Fatal("Result must be Hello Abik") // jika gagal kode programnya di hentikan atau memanggil FailNow
	}
	fmt.Println("TestHelloWorldFatal done")
}

// testing menggunakan Testify
// dengan kata kunci assert ( dengan assert akan memanggi fungsi Fail)
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld(" Abik")
	assert.Equal(t, "Hello Abik", result, "Result must be Hello Abik")
	fmt.Println("Test with assert")
}

// dengan kata kunci require ( dengan require akan memanggi fungsi FailNow)
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld(" Abik")
	assert.Equal(t, "Hello Abik", result, "Result must be Hello Abik")
	fmt.Println("Test with require")
}

// Menggunakan SkipTest yaitu jika kondisi terpenuhi, maka test dilewati

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak jalan di Mac")
	}
	result := HelloWorld("masendhy")
	require.Equal(t, "Hello masendhy", result, "Result maus be Hello masendhy")
}

//Before dan After Test
func TestMain(m *testing.M) {
	// before test
	fmt.Println("before unit test")

	m.Run()

	//after test
	fmt.Println("After test")
}

// SubTest : testing di dalam testing
func TestSubTest(t *testing.T) {
	t.Run("masendhy", func(t *testing.T) {
		result := HelloWorld(" masendhy")
		require.Equal(t, "Hello masendhy", result, "Result must be Hello masendhy")
	})

	t.Run("abik", func(t *testing.T) {
		result := HelloWorld(" abik")
		require.Equal(t, "Hello abik", result, "Result must be Hello abik")
	})
}

//TableTest

func TestTableHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "masendhy",
			request:  " masendhy",
			expected: "Hello masendhy",
		},
		{
			name:     "abik",
			request:  " abik",
			expected: "Hello abik",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}
