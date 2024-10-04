// Package smallestpositive menyediakan fungsi untuk menemukan
// bilangan bulat positif terkecil yang tidak terdapat dalam array.
package smallestpositive

// SmallestPositive mencari bilangan bulat positif terkecil yang
// tidak ada dalam array nums. Fungsi ini mengembalikan angka tersebut.
func SmallestPositive(nums []int) int {
	// Membuat map untuk menyimpan bilangan positif yang ditemukan
	numSet := make(map[int]struct{})

	// Mengisi map dengan bilangan positif dari array
	for _, num := range nums {
		if num > 0 {
			numSet[num] = struct{}{}
		}
	}

	// Memulai pencarian dari bilangan positif terkecil, yaitu 1
	i := 1
	for {
		// Mengecek apakah i ada dalam numSet
		if _, exists := numSet[i]; !exists {
			// Jika tidak ada, kembalikan nilai i sebagai hasil
			return i
		}
		i++ // Jika ada, lanjutkan ke bilangan berikutnya
	}
}
