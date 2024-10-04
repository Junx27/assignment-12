package inttoroman //deklarasi package

func IntToRoman(num int) string {
	// Mendefinisikan slice dari struktur yang menyimpan nilai dan simbol Romawi
	romanNumerals := []struct {
		value  int    // nilai integer
		symbol string // simbol Romawi yang sesuai
	}{
		{1000, "M"},
		{500, "D"},
		{100, "C"},
		{50, "L"},
		{10, "X"},
		{5, "V"},
		{1, "I"},
	}

	result := ""                       // Inisialisasi string kosong untuk menyimpan hasil konversi
	for _, rn := range romanNumerals { // Iterasi melalui setiap nilai dan simbol Romawi

		// Selama angka yang tersisa lebih besar dari nilai saat ini
		for num > rn.value {
			switch { // Memeriksa kondisi untuk menentukan simbol yang harus ditambahkan
			case num >= 1000:
				result += rn.symbol // Tambahkan simbol ke hasil
				num -= rn.value     // Kurangi nilai dari num
			case num >= 900:
				result += "CM" // Tambahkan simbol "CM"
				num -= 900     // Kurangi 900 dari num
			case num >= 500:
				result += rn.symbol // Tambahkan simbol ke hasil
				num -= rn.value     // Kurangi nilai dari num
			case num >= 400:
				result += "CD" // Tambahkan simbol "CD"
				num -= 400     // Kurangi 400 dari num
			case num >= 100:
				result += rn.symbol // Tambahkan simbol ke hasil
				num -= rn.value     // Kurangi nilai dari num
			case num >= 90:
				result += "XC" // Tambahkan simbol "XC"
				num -= 90      // Kurangi 90 dari num
			case num >= 50:
				result += rn.symbol // Tambahkan simbol ke hasil
				num -= rn.value     // Kurangi nilai dari num
			case num >= 40:
				result += "XL" // Tambahkan simbol "XL"
				num -= 40      // Kurangi 40 dari num
			case num >= 10:
				result += rn.symbol // Tambahkan simbol ke hasil
				num -= rn.value     // Kurangi nilai dari num
			case num >= 9:
				result += "IX" // Tambahkan simbol "IX"
				num -= 9       // Kurangi 9 dari num
			case num >= 5:
				result += rn.symbol // Tambahkan simbol ke hasil
				num -= rn.value     // Kurangi nilai dari num
			case num >= 4:
				result += "IV" // Tambahkan simbol "IV"
				num -= 4       // Kurangi 4 dari num
			case num >= 1:
				result += rn.symbol // Tambahkan simbol ke hasil
				num -= rn.value     // Kurangi nilai dari num
			}
		}
	}

	return result // Kembalikan hasil konversi sebagai string
}
