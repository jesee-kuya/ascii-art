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

func TestAscii(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			content, _ := ascii.Reader("/home/jkuya/ascii-art/standard.txt", "\n")
			input := strings.Split(tc.input, "\n")

			ascii.Ascii(content, input)

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			_, err := buf.ReadFrom(r)
			if err != nil {
				t.Fatalf("Error reading from pipe: %v", err)
			}
			expected := tc.line1 + tc.line2 + tc.line3 + tc.line4 + tc.line5 + tc.line6 + tc.line7 + tc.line8
			if buf.String() != expected {
				t.Errorf("got \n %v expected \n %v", buf.String(), expected)
			}
		})
	}
}
