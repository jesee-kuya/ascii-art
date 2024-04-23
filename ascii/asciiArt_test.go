package ascii_test

import (
	"ascii/ascii"
	"os"
	"testing"
)

func TestAscii(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	// defer func() {
	os.Stdout = old
	// }
	content := ascii.Reader("standard.txt", "\n")
	input := []string {"Hello"}
	ascii.Ascii(content, input)

	w.Close()

	out := make([]byte, 1024)
	n, _ := r.Read(out)
	actualOutput := string(out[:n])
	
	expectedOutput := ` 
	 _    _          _   _          
	| |  | |        | | | |         
	| |__| |   ___  | | | |   ___   
	|  __  |  / _ \ | | | |  / _ \  
	| |  | | |  __/ | | | | | (_) | 
	|_|  |_|  \___| |_| |_|  \___/  
									
									
	`

	if actualOutput != expectedOutput {
		t.Errorf("got %v, \n expected %v", actualOutput, expectedOutput)
	}
}

    
