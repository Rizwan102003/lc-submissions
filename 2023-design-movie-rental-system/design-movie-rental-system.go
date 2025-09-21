// [ price (16 bits) | shop (32 bits) | movie (15 bits) | rented (1 bit) ]
func pack(shop, movie, price int) int {
	return (price << 49) | (shop << 17) | (movie << 1)
}

func unpack(x int) (shop, movie, price, rented int) {
	rented = x & 1
	movie = (x >> 1) & 0xFFFF
	shop = (x >> 17) & 0xFFFFFFFF
	price = x >> 49
	return
}

type Price = int

var ENTRIES [100000]int
var INDEXMAP = make(map[[2]int]int, 1000)

type MovieRentingSystem struct {
	cheapestUnrentedByMovies map[int][]int
	cheapestRented           []int
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	clear(INDEXMAP)
	movies := map[int][]int{}
	for i, entry := range entries {
		price := pack(entry[0], entry[1], entry[2])
		movies[entry[1]] = append(movies[entry[1]], price)
		ENTRIES[i] = price
		INDEXMAP[[2]int{entry[0], entry[1]}] = i
	}
	for movie := range movies {
		slices.Sort(movies[movie])
	}
	return MovieRentingSystem{
		cheapestUnrentedByMovies: movies,
		cheapestRented:           make([]int, 0, 16),
	}
}

func (mvr *MovieRentingSystem) Search(movie int) []int {
	cheapestUnrented := make([]int, 0, 5)
	unrented := mvr.cheapestUnrentedByMovies[movie]
	for i := 0; i < len(unrented) && len(cheapestUnrented) < 5; i++ {
		shop, _, _, _ := unpack(unrented[i])
		idx := INDEXMAP[[2]int{shop, movie}]
		if ENTRIES[idx]&1 == 0 {
			cheapestUnrented = append(cheapestUnrented, shop)
		}
	}
	return cheapestUnrented
}

func (mvr *MovieRentingSystem) Rent(shop int, movie int) {
	idx := INDEXMAP[[2]int{shop, movie}]
	ENTRIES[idx] |= 1
	mvr.cheapestRented = append(mvr.cheapestRented, ENTRIES[idx])
	slices.Sort(mvr.cheapestRented)
}

func (mvr *MovieRentingSystem) Drop(shop int, movie int) {
	idx := INDEXMAP[[2]int{shop, movie}]
	i := slices.Index(mvr.cheapestRented, ENTRIES[idx])
	n := len(mvr.cheapestRented) - 1
	mvr.cheapestRented[i], mvr.cheapestRented[n] = mvr.cheapestRented[n], mvr.cheapestRented[i]
	mvr.cheapestRented = mvr.cheapestRented[0:n]
	ENTRIES[idx] &^= 1
	slices.Sort(mvr.cheapestRented)
}

func (mvr *MovieRentingSystem) Report() [][]int {
	n := min(5, len(mvr.cheapestRented))
	report := make([][]int, n)
	for i := range n {
		shop, movie, _, _ := unpack(mvr.cheapestRented[i])
		report[i] = []int{shop, movie}
	}
	return report
}

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */