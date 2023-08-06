package testtaker

import (
	"sort"
)

func MaxPleaseure(n int, k int, size []int, taste []int) int {
	if n == k {
		return -1
	}
	
	type Cake struct {
		size, taste int
	}
	
	cakes := make([]Cake, n)
	for i := 0; i < n; i++ {
		cakes[i] = Cake{size[i], taste[i]}
	}
	
	sort.Slice(cakes, func(i, j int) bool {
		return cakes[i].taste > cakes[j].taste
	})
	
	bestSetTaste := 0
	for i := 0; i < k; i++ {
		bestSetTaste = max(bestSetTaste, cakes[i].taste)
	}
	
	pleasure := 0
	for i := 0; i < k; i++ {
		pleasure += cakes[i].size * bestSetTaste
	}
	
	return pleasure
}


func main() {
	n := 5
	k := 3
	sizes := []int{2, 4, 1, 5, 3}
	tastes := []int{8, 5, 6, 9, 7}

	maxPleasure := testtaker.MaxPleaseure(n, k, sizes, tastes)

	fmt.Printf("Kesenangan maksimal yang bisa Anda dapatkan adalah %d\n", maxPleasure)
}