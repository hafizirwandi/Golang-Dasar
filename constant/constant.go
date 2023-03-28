/*
- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya
*/

package main

import "fmt"

func main(){
	const phi float32 = 3.14
	fmt.Println("Konstanta Phi adalah ",phi)
}