package main

import (
	"fmt"
)

func main() {
	var berat, tinggiCm float64
	var bmi float64
	var kategori string

	fmt.Println("==========================")
	fmt.Println("===== Kalkulator BMI =====")
	fmt.Println("==========================")
	fmt.Print("\nMasukkan berat badan (kg): ")
	_, err := fmt.Scan(&berat)

	if err != nil {
		fmt.Println("Tipe data invalid. Masukkan angka untuk berat badan. Contoh input yang benar: 70.5")
		return
	}
	if berat <= 0 || berat >= 300 {
		fmt.Println("Berat badan harus lebih dari 0 kg dan tidak boleh lebih dari 300 kg. Contoh input yang benar: 70")
		return
	}

	fmt.Print("Masukkan tinggi badan (cm): ")
	_, err = fmt.Scan(&tinggiCm)

	if err != nil {
		fmt.Println("Tipe data invalid. Masukkan angka untuk tinggi badan. Contoh input yang benar: 170.5")
		return
	}
	if tinggiCm <= 0 || tinggiCm > 300 {
		fmt.Println("Tinggi badan harus lebih dari 0 cm dan tidak boleh lebih ari 300 cm. Contoh input yang benar: 170")
		return
	}

	tinggiMeter := tinggiCm / 100
	bmi = berat / (tinggiMeter * tinggiMeter)

	if bmi < 18.5 {
		kategori = "Kurus"
	} else if bmi >= 18.5 && bmi <= 24.9 {
		kategori = "Normal"
	} else if bmi >= 25 && bmi <= 29.9 {
		kategori = "Gemuk"
	} else {
		kategori = "Obesitas"
	}

	fmt.Printf("\nHasil Perhitungan BMI:\n")
	fmt.Printf("Berat badan: %.2f kg\n", berat)
	fmt.Printf("Tinggi badan: %.2f cm\n", tinggiCm)
	fmt.Printf("Nilai BMI: %.2f\n", bmi)
	fmt.Printf("Kategori: %s\n", kategori)

	fmt.Println("\nKeterangan BMI:")
	fmt.Println("Kurus    (BMI < 18.5)   : Perlu meningkatkan berat badan")
	fmt.Println("Normal   (18.5-24.9)    : Berat badan ideal")
	fmt.Println("Gemuk    (25-29.9)      : Perlu menurunkan berat badan")
	fmt.Println("Obesitas (BMI ≥ 30)     : Perlu konsultasi dokter")

	fmt.Printf("\n\n !!! Hasil BMI tidak efektif untuk usia kurang dari 18 tahun dan Ibu Hamil !!! \n\n")
}