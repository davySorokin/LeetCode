package main

import (
	"container/heap"
)

type Food struct {
	name   string
	rating int
}

// Implement heap.Interface for FoodHeap
type FoodHeap []*Food

func (h FoodHeap) Len() int { return len(h) }

func (h FoodHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].name < h[j].name // Lexicographically smaller comes first if tied
	}
	return h[i].rating > h[j].rating // Higher rating comes first
}

func (h FoodHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push/Pop for heap.Interface
func (h *FoodHeap) Push(x interface{}) {
	*h = append(*h, x.(*Food))
}

func (h *FoodHeap) Pop() interface{} {
	old := *h
	n := len(old)
	food := old[n-1]
	*h = old[0 : n-1]
	return food
}

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRating  map[string]int
	cuisineToHeap map[string]*FoodHeap
}

// Constructor
func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineToHeap: make(map[string]*FoodHeap),
	}

	for i := range foods {
		food := foods[i]
		cuisine := cuisines[i]
		rating := ratings[i]

		fr.foodToCuisine[food] = cuisine
		fr.foodToRating[food] = rating

		if _, exists := fr.cuisineToHeap[cuisine]; !exists {
			fr.cuisineToHeap[cuisine] = &FoodHeap{}
		}
		heap.Push(fr.cuisineToHeap[cuisine], &Food{name: food, rating: rating})
	}

	return fr
}

// ChangeRating updates a food's rating
func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine := fr.foodToCuisine[food]
	fr.foodToRating[food] = newRating

	// Push new version of the food onto the heap
	heap.Push(fr.cuisineToHeap[cuisine], &Food{name: food, rating: newRating})
}

// HighestRated returns the top-rated food for a given cuisine
func (fr *FoodRatings) HighestRated(cuisine string) string {
	h := fr.cuisineToHeap[cuisine]
	for {
		top := (*h)[0]
		currentRating := fr.foodToRating[top.name]
		if top.rating == currentRating {
			return top.name
		}
		// Stale entry — remove and check next
		heap.Pop(h)
	}
}
