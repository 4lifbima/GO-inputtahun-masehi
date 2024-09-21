package main

import (
	"fmt"
	"time"
)

func main() {
	var month, year int

	// Input tahun
	fmt.Print("Masukkan tahun (misal: 2024): ")
	fmt.Scan(&year)

	// Input bulan
	fmt.Print("Masukkan bulan (1-12): ")
	fmt.Scan(&month)

	// Cek apakah bulan valid
	if month < 1 || month > 12 {
		fmt.Println("Bulan tidak valid. Masukkan angka antara 1 dan 12.")
		return
	}

	// Membuat objek time untuk hari pertama bulan yang diinput
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)

	// Menemukan berapa hari dalam bulan tersebut
	daysInMonth := time.Date(year, firstDay.Month()+1, 0, 0, 0, 0, 0, time.Local).Day()

	// Menampilkan kalender sederhana
	fmt.Printf("\nKalender untuk %s %d\n", firstDay.Month(), year)
	fmt.Println("Su Mo Tu We Th Fr Sa")

	// Menemukan hari pertama bulan itu (0=Sunday, 1=Monday, ...)
	startDay := int(firstDay.Weekday())

	// Print space untuk hari sebelum tanggal 1
	for i := 0; i < startDay; i++ {
		fmt.Print("   ")
	}

	// Print tanggal-tanggal dalam bulan
	for day := 1; day <= daysInMonth; day++ {
		fmt.Printf("%2d ", day)
		if (day+startDay)%7 == 0 { // Pindah ke baris berikutnya setiap Sabtu
			fmt.Println()
		}
	}

	fmt.Println()
}
