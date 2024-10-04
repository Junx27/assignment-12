// Package twosum menyediakan fungsi untuk menemukan dua indeks
// dalam array yang jumlahnya sama dengan target tertentu.
package twosum

// TwoSum mencari dua angka dalam array nums yang jika dijumlahkan
// menghasilkan nilai target. Fungsi ini mengembalikan indeks dari
// dua angka tersebut sebagai slice dari int.
func TwoSum(nums []int, target int) []int {
	// Membuat map untuk menyimpan nilai dan indeksnya
	numMap := make(map[int]int)

	// Melakukan iterasi pada setiap elemen dalam array nums
	for i, num := range nums {
		// Menghitung nilai yang diperlukan untuk mencapai target
		complement := target - num

		// Mengecek apakah complement sudah ada dalam map
		if idx, found := numMap[complement]; found {
			// Jika ditemukan, kembalikan indeks dari complement dan indeks saat ini
			return []int{idx, i}
		}

		// Jika tidak ditemukan, simpan nilai dan indeks saat ini dalam map
		numMap[num] = i
	}

	// Jika tidak ada pasangan yang ditemukan, kembalikan nil
	return nil
}
