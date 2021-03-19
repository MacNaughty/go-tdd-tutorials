package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t1 *testing.T) {
	assertCorrectMessage := func(t2 testing.TB, got string, want string) {
		t2.Helper()
		if got != want {
			t2.Errorf("got %q want %q", got, want)
		}
	}

	t1.Run("repeat character a 5 times", func(t2 *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t2, repeated, expected)
	})

	t1.Run("repeat character a 5 times", func(t2 *testing.T) {
		repeated := Repeat("b", 9)
		expected := "bbbbbbbbb"
		assertCorrectMessage(t2, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeatedChar := Repeat("a", 16)
	fmt.Println(repeatedChar)
	// Output: "aaaaaaaaaaaaaaaa"
}
