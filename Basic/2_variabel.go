package main

import "fmt"

func main() {
	var nama string //deklarasi variabel
	nama = "Rizq"
	fmt.Println(nama)

	var nama_awal = "Muhammad"

	nama_akhir := "Ramadhan"
	fmt.Println("nama akhir:", nama_akhir)

	fmt.Println("Nama saya:", nama_awal, nama, nama_akhir)
	fmt.Printf("type nama adalah %T", nama)
}
