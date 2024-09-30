import pandas as pd

def count_followers(followers: pd.DataFrame) -> pd.DataFrame:
    # Group by user_id and count the number of followers for each user
    follower_counts = followers.groupby('user_id').size().reset_index(name='followers_count')
    
    # Sort the result by user_id in ascending order
    result = follower_counts.sort_values('user_id')
    
    return result
