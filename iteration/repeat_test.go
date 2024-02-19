package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeating a character by using string library function", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectRepeat(t, repeated, expected)
	})

	t.Run("fail", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaa"
		assertCorrectRepeat(t, repeated, expected)
	})
}

func TestRepeatIterable(t *testing.T) {
	t.Run("repeating a character using iterable", func(t *testing.T) {
		repeated := RepeatIterable("a", 5)
		expected := "aaaaa"
		assertCorrectRepeat(t, repeated, expected)
	})

	t.Run("fail", func(t *testing.T) {
		repeated := RepeatIterable("a", 6)
		expected := "aaaaa"
		assertCorrectRepeat(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

func assertCorrectRepeat(t testing.TB, expected, repeated string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
