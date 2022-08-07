package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAggregate(t *testing.T) {
	got := Aggregate()
	want := `    2016-12-20T19:00:45Z, Server A started.
	2016-12-20T19:01:16Z, Server B started.
	2016-12-20T19:01:25Z, Server A completed job.
	2016-12-20T19:02:48Z, Server A terminated.
	2016-12-20T19:03:25Z, Server B completed job.
	2016-12-20T19:04:50Z, Server B terminated.`

	if got != want {
		t.Errorf("Unexpected output. got: %s, want: %s", got, want)
	}
}
