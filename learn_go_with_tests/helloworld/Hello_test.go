package helloworld

import (
	"testing"

)

func TestHello(t *testing.T) {

	t.Run("Saying Hello to People", func(t *testing.T) {
		got := Hello("Vinay","")
		want := "Hello, Vinay"

		Assert(t, want, got)
	})
    t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"

		Assert(t, want, got)
	})
    t.Run("Saying Hello in Spanish", func(t *testing.T) {
        got:= Hello("Aman","Spanish")
        want := "Hola, Aman"

        Assert(t, want, got)
    })

    t.Run("Saying Hello in French", func(t *testing.T) {
        got:= Hello("Sonali","French")
        want := "Bonjour, Sonali"

        Assert(t, want, got)
    })

    t.Run("Saying Hello in Hindi", func(t *testing.T) {
        got:= Hello("Sanjay","Hindi")
        want := "Namaste, Sanjay"

        Assert(t, want, got)
    })
}

func Assert(t testing.TB, want, got string) {
    t.Helper()
    if want != got {
        t.Errorf("Expected %s but got %v", want, got)
    }


}
