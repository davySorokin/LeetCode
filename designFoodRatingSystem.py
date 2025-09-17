from collections import defaultdict
import bisect

class FoodRatings:
    def __init__(self, foods, cuisines, ratings):
        self.food_map = {}
        self.cuisine_map = defaultdict(list)

        for food, cuisine, rating in zip(foods, cuisines, ratings):
            self.food_map[food] = {'rating': rating, 'cuisine': cuisine}
            bisect.insort(self.cuisine_map[cuisine], (-rating, food))

    def changeRating(self, food, newRating):
        # Update the rating
        old_rating = self.food_map[food]['rating']
        cuisine = self.food_map[food]['cuisine']
        self.food_map[food]['rating'] = newRating
        
        # Remove old rating from the cuisine list
        self.cuisine_map[cuisine].remove((-old_rating, food))
        bisect.insort(self.cuisine_map[cuisine], (-newRating, food))

    def highestRated(self, cuisine):
        # The highest-rated food is the first in the sorted list (most negative rating)
        return self.cuisine_map[cuisine][0][1]


# usage
foodRatings = FoodRatings(["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], 
                          ["korean", "japanese", "japanese", "greek", "japanese", "korean"], 
                          [9, 12, 8, 15, 14, 7])

print(foodRatings.highestRated("korean"))   # "kimchi"
print(foodRatings.highestRated("japanese")) # "ramen"
foodRatings.changeRating("sushi", 16)
print(foodRatings.highestRated("japanese")) # "sushi"
foodRatings.changeRating("ramen", 16)
print(foodRatings.highestRated("japanese")) # "ramen"
