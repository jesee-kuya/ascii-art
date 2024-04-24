package ascii_test

import (
	"bytes"
	"os"
	"testing"

	"ascii/ascii"
)

func TestAscii(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	content := ascii.Reader("/home/jkuya/ascii-art/standard.txt", "\n")
	input := []string{"Hello"}

	ascii.Ascii(content, input)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Error reading from pipe: %v", err)
	}
	line1 := " _    _          _   _          \n"
	line2 := "| |  | |        | | | |         \n"
	line3 := "| |__| |   ___  | | | |   ___   \n"
	line4 := "|  __  |  / _ \\ | | | |  / _ \\  \n"
	line5 := "| |  | | |  __/ | | | | | (_) | \n"
	line6 := "|_|  |_|  \\___| |_| |_|  \\___/  \n"
	line7 := "                                \n"
	line8 := "                                \n"
	expected := line1 + line2 + line3 + line4 + line5 + line6 + line7 + line8
	if buf.String() != expected {
		t.Errorf("got \n %v expected \n %v", buf.String(), expected)
	}
}
