import pandas as pd

round2 = lambda x: round(x + 1e-9, 2)  # Rounds to avoid floating-point precision issues

def queries_stats(queries: pd.DataFrame) -> pd.DataFrame:
    # Calculate quality as rating divided by position
    queries['quality'] = queries['rating'] / queries['position']
    
    # Calculate poor_query_percentage as (rating < 3) multiplied by 100
    queries['poor_query_percentage'] = (queries['rating'] < 3) * 100
    
    # Group by query_name and calculate the mean of quality and poor_query_percentage
    result = queries.groupby('query_name')[['quality', 'poor_query_percentage']].mean()
    
    # Apply the rounding function to both columns
    result = result.apply(round2).reset_index()
    
    return result
