package main

import ascii "ascii/ascii"

func ExampleAscii() {
	file := ascii.Reader("standard.txt", "\n")
	input := []string{"Hello"}
	ascii.Ascii(file, input)

	// Unordered output:
	//   _    _          _   _          
	//  | |  | |        | | | |         
	//  | |__| |   ___  | | | |   ___   
	//  |  __  |  / _ \ | | | |  / _ \  
	//  | |  | | |  __/ | | | | | (_) | 
	//  |_|  |_|  \___| |_| |_|  \___/  
	//  								
	//	    							
	//
}
