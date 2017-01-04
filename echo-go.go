package main

import "os"

func main() {
	b := make([]byte, 64)
	i := 0
	for n, err := os.Stdin.Read(b); err == nil && n > 0; n, err = os.Stdin.Read(b) {
		i += 1
		os.Stdout.Write(b[0:n])
	}
}
