package command

import (
	"fmt"
	"testing"
)

func TestFindCommand(t *testing.T) {
	r := NewRoute("?")
	r.Cmd = make(map[string]*Command)

	for i := 0; i < 100; i++ {
		r.Cmd[fmt.Sprintf("%d", i)] = &Command{Name: fmt.Sprintf("%d", i)}
	}

	want := "99"

	got := r.FindCommand("99")

	if got.Name != want {
		t.Error(" got: ", got.Name, "; Needed: ", want)
	}
}

func BenchmarkFindCommand(b *testing.B) {
	r := NewRoute("?")
	r.Cmd = make(map[string]*Command)

	for n := 1; n <= 1000000; n *= 10 {
		b.Run(fmt.Sprintf("CommandsRegistered : %d", n), func(b *testing.B) {
			for i := 1; i <= n; i++ {
				r.Cmd[fmt.Sprintf("%d", i)] = &Command{Name: fmt.Sprintf("%d", i)}
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				r.FindCommand(fmt.Sprintf("%d", n))
			}
		})
	}
}
