package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a character 5 times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat a character a custom number of times", func(t *testing.T) {
		customRepeatCount := 3
		repeated := RepeatCustom("b", customRepeatCount)
		expected := "bbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a"))
	// Output: aaaaa
}
