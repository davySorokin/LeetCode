import pandas as pd

def not_boring_movies(cinema: pd.DataFrame) -> pd.DataFrame:
    # Filter movies with odd-numbered IDs and description not 'boring'
    filtered_movies = cinema[(cinema['id'] % 2 != 0) & (cinema['description'] != 'boring')]
    
    # Sort the filtered movies by rating in descending order
    sorted_movies = filtered_movies.sort_values(by='rating', ascending=False)
    
    # Return the relevant columns
    return sorted_movies[['id', 'movie', 'description', 'rating']]
