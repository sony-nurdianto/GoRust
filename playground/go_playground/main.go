package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic tertangkap:", r)
		}
	}()

	fmt.Println("Sebelum panic")
	panic("Terjadi kesalahan!")  // Panic terjadi di sini
	fmt.Println("Setelah panic") // Tidak akan dieksekusi
}
