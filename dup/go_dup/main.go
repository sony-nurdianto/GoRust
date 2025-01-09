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
	// Membuat HashMap untuk menyimpan jumlah kemunculan setiap baris.
	args_len := len(os.Args[1:])
	fmt.Println(args_len)

	counts := make(map[string]int)

	// Membaca input dari stdin menggunakan scanner
	input := bufio.NewScanner(os.Stdin)

	// Melakukan iterasi pada setiap baris input, jika input valid maka akan memperbarui nilai pada hashmap
	for input.Scan() {
		counts[input.Text()]++
	}

	// Menampilkan hasil dari baris yang memiliki nilai duplikat
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
