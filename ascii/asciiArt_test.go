package ascii_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"ascii/ascii"
)

var testCases = []struct {
	name                                                   string
	line1, line2, line3, line4, line5, line6, line7, line8 string
	input                                                  string
}{
	{
		"Test1",
		" _    _          _   _          \n",
		"| |  | |        | | | |         \n",
		"| |__| |   ___  | | | |   ___   \n",
		"|  __  |  / _ \\ | | | |  / _ \\  \n",
		"| |  | | |  __/ | | | | | (_) | \n",
		"|_|  |_|  \\___| |_| |_|  \\___/  \n",
		"                                \n",
		"                                \n",
		"Hello",
	},
	{
		"Test2",
		" _    _          _   _          \n",
		"| |  | |        | | | |         \n",
		"| |__| |   ___  | | | |   ___   \n",
		"|  __  |  / _ \\ | | | |  / _ \\  \n",
		"| |  | | |  __/ | | | | | (_) | \n",
		"|_|  |_|  \\___| |_| |_|  \\___/  \n",
		"                                \n",
		"                                \n\n",
		"Hello\n",
	},
}

// TestAscii tests the Ascii function
func TestAscii(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// redirect stdout to pipe
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// read ascii art from a file
			content, _ := ascii.Reader("/home/jkuya/ascii-art/standard.txt", "\n")
			// split input into lines
			input := strings.Split(tc.input, "\n")

			// run the ascii function
			ascii.Ascii(content, input)

			// close the write end of pipe
			w.Close()
			// restore stdout
			os.Stdout = old

			// read from the read end of pipe
			var buf bytes.Buffer
			_, err := buf.ReadFrom(r)
			if err != nil {
				t.Fatalf("Error reading from pipe: %v", err)
			}
			// compare output to expected output
			expected := tc.line1 + tc.line2 + tc.line3 + tc.line4 + tc.line5 + tc.line6 + tc.line7 + tc.line8
			if buf.String() != expected {
				t.Errorf("got \n %v expected \n %v", buf.String(), expected)
			}
		})
	}
}
