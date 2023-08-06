package testtaker

import (
	"sort"
)

func MaxPleaseure(n int, k int, size []int, taste []int) int {
	if n == k {
		return -1
	}

	// membuat slice kue dengan ukuran dan rasa yang sesuai
	var cakes = []Cake
	for i := 0; i < n; i++ {
		cakes = append(cakes, Cake{Size: size[i], Taste: taste[i]})
	}

	// mengurutkan kue berdasarkan rasa
	sort.Slice(cakes, func(i, j int) bool {
		return cakes[i].Taste < cakes[j].Taste
	})

	// memilih kue dengan rasa terendah dan ukuran terbesar
	var totalSize int
	var minTaste int = cakes[0].Taste
	for i := 0; i < k; i++ {
		if cakes[i].Size > cakes[n-1-i].Size {
			totalSize += cakes[i].Size
		} else {
			totalSize += cakes[n-1-i].Size
		}
	}

	// menghitung kesenangan maksimum yang bisa didapatkan
	maxPleaseure := totalSize * minTaste

	return maxPleaseure
}

type Cake struct {
	Size  int
	Taste int
}