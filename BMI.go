package main

import "fmt"

func main() {
	for {
		var berat, tinggiCm float64
		var bmi float64
		var kategori string

		fmt.Println("==========================")
		fmt.Println("===== Kalkulator BMI =====")
		fmt.Println("==========================")

		for {
			fmt.Print("\nMasukkan berat badan (kg): ")
			_, err := fmt.Scan(&berat)

			if err != nil {
				var ngasal string
				fmt.Scan(&ngasal)
				fmt.Println("Input harus berupa angka. Contoh: 70.5")
				continue
			}

			if berat <= 0 || berat >= 300 {
				fmt.Println("Berat badan harus antara 0 - 300 kg.")
				continue
			}

			break
		}

		for {
			fmt.Print("Masukkan tinggi badan (cm): ")
			_, err := fmt.Scan(&tinggiCm)

			if err != nil {
				var ngasal string
				fmt.Scan(&ngasal)
				fmt.Println("Input harus berupa angka. Contoh: 170cm")
				continue
			}

			if tinggiCm <= 0 || tinggiCm > 300 {
				fmt.Println("Tinggi badan harus antara 0 - 300 cm.")
				continue
			}

			break
		}

		tinggiMeter := tinggiCm / 100
		bmi = berat / (tinggiMeter * tinggiMeter)

		if bmi < 18.5 {
			kategori = "Kurus"
		} else if bmi <= 24.9 {
			kategori = "Normal"
		} else if bmi <= 29.9 {
			kategori = "Gemuk"
		} else {
			kategori = "Obesitas"
		}

		fmt.Println("\nHasil Perhitungan BMI")
		fmt.Printf("Berat badan  : %.2f kg\n", berat)
		fmt.Printf("Tinggi badan : %.2f cm\n", tinggiCm)
		fmt.Printf("Nilai BMI    : %.2f\n", bmi)
		fmt.Printf("Kategori     : %s\n", kategori)

		fmt.Println("\n!!! Hasil BMI tidak efektif untuk usia < 18 tahun dan ibu hamil !!!")

		var ulang string
		fmt.Print("\nApakah ingin mengulangi perhitungan (ya/tidak)? ")
		fmt.Scan(&ulang)

		if ulang != "ya" && ulang != "y" {
			fmt.Println("\nProgram selesai")
			break
		}
	}
}
