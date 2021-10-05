package main

import (
	"fmt"
	"os"
)

type memberGo struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	var coba memberGo
	//var err validation

	masukkanMenu := os.Args[1]

	if masukkanMenu == "GLNG016ONL001" {
		coba.Nama = "Mas Dendhy Nugroho"
		coba.Alamat = "Jakarta"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Programer"
	} else if masukkanMenu == "GLNG016ONL002" {
		coba.Nama = "Haris Ardianto Wibowo"
		coba.Alamat = "Jakarta"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Software Engineer"
	} else if masukkanMenu == "GLNG016ONL003" {
		coba.Nama = "Hidayat Aji Santoso"
		coba.Alamat = "Depok"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Software Engineer"
	} else if masukkanMenu == "GLNG016ONL004" {
		coba.Nama = "Muhammad Alif Pratama"
		coba.Alamat = "Jakarta"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Programer"
	} else if masukkanMenu == "GLNG016ONL005" {
		coba.Nama = "Muhammmad Reyhan"
		coba.Alamat = "Bali"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Programer"
	} else if masukkanMenu == "GLNG016ONL006" {
		coba.Nama = "Imam Setia Utama N"
		coba.Alamat = "Bogor"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Software Engineer"
	} else if masukkanMenu == "GLNG016ONL007" {
		coba.Nama = "Razan Latang Rahardjo"
		coba.Alamat = "Sukabumi"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Programer"
	} else if masukkanMenu == "GLNG016ONL008" {
		coba.Nama = "Mohtar Nurwahid"
		coba.Alamat = "Garut"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Software Engineer"
	} else if masukkanMenu == "GLNG016ONL009" {
		coba.Nama = "Dwi Ahmad Hisyam"
		coba.Alamat = "Manado"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Software Engineer"
	} else if masukkanMenu == "GLNG016ONL010" {
		coba.Nama = "Barikly Mubarok"
		coba.Alamat = "Kanada"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Programmer"
	} else if masukkanMenu == "GLNG016ONL011" {
		coba.Nama = "Seven"
		coba.Alamat = "Jepang"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Software Engineer"
	} else if masukkanMenu == "GLNG016ONL012" {
		coba.Nama = "Arnold"
		coba.Alamat = "Amerika"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Software Engineer"
	} else if masukkanMenu == "GLNG016ONL013" {
		coba.Nama = "Beni"
		coba.Alamat = "Indonesia"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Software Engineer"
	} else if masukkanMenu == "GLNG016ONL014" {
		coba.Nama = "Ivan"
		coba.Alamat = "Qatar"
		coba.Alasan = "Belajar Golang"
		coba.Pekerjaan = "Software Engineer"
	} else {
		fmt.Println("Kode Data Member Tidak ditemukan")
	}

	fmt.Println("Nama: ", coba.Nama)
	fmt.Println("Alamat: ", coba.Alamat)
	fmt.Println("Alasan: ", coba.Alasan)
	fmt.Println("Pekerjaan: ", coba.Pekerjaan)
}
