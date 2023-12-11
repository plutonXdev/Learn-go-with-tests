package main

import "testing"

func TestHello(t *testing.T)  {
	 got :=Hello("Ace")

	 want := "Hello,Ace"

	 if got !=want{
          t.Errorf("got %q want %q",got,want)
	 }

	 }
