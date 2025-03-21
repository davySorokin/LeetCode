from collections import deque, defaultdict
from typing import List

class Solution:
    def findAllRecipes(self, recipes: List[str], ingredients: List[List[str]], supplies: List[str]) -> List[str]:
        # Initialize adjacency list and in-degree count
        indegree = {}
        graph = defaultdict(list)
        supply_set = set(supplies)
        
        # Build graph and in-degree map
        for recipe, ing_list in zip(recipes, ingredients):
            indegree[recipe] = len(ing_list)  # Number of missing ingredients
            for ing in ing_list:
                graph[ing].append(recipe)  # Ingredient is required by this recipe
        
        # Start with all supplies in the queue
        queue = deque(supplies)
        result = []

        # Process the queue
        while queue:
            item = queue.popleft()
            
            # If the item is a recipe, we can make it
            if item in indegree:
                result.append(item)

            # Check recipes that depend on this ingredient
            for recipe in graph[item]:
                if recipe in indegree:
                    indegree[recipe] -= 1  # Reduce missing ingredient count
                    if indegree[recipe] == 0:  # If all ingredients are available
                        queue.append(recipe)

        return result
