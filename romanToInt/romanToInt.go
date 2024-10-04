// Package romantoint menyediakan fungsi untuk mengonversi
// angka romawi menjadi integer.
package romantoint

// RomanToInt mengonversi string angka romawi menjadi integer.
// Fungsi ini mengembalikan nilai integer dari string romawi yang diberikan.
func RomanToInt(s string) int {
	// Membuat map untuk menyimpan nilai dari setiap simbol romawi
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0     // Variabel untuk menyimpan total nilai
	prevValue := 0 // Variabel untuk menyimpan nilai sebelumnya

	// Mengiterasi string dari belakang ke depan
	for i := len(s) - 1; i >= 0; i-- {
		value := romanMap[s[i]] // Mendapatkan nilai dari simbol romawi
		if value < prevValue {  // Jika nilai saat ini lebih kecil dari nilai sebelumnya
			total -= value // Kurangi total dengan nilai saat ini
		} else {
			total += value // Tambahkan nilai saat ini ke total
		}
		prevValue = value // Update nilai sebelumnya
	}
	return total // Kembalikan total nilai
}
