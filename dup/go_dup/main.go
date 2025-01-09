// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Program utama untuk menghitung frekuensi kemunculan setiap baris input
// yang dibaca dari stdin.
//
// Program ini akan membaca setiap baris yang dimasukkan, menghitung
// jumlah kemunculan setiap baris, dan kemudian mencetak hasilnya dalam
// bentuk `HashMap` di mana kunci adalah baris input dan nilai adalah
// jumlah kemunculannya.

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
