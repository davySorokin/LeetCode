import pandas as pd

def user_activity(activity: pd.DataFrame) -> pd.DataFrame:
    # Filter activity data for the 30 days ending on 2019-07-27
    filtered_activity = activity[(activity['activity_date'] >= '2019-06-28') & (activity['activity_date'] <= '2019-07-27')]
    
    # Group by activity_date and user_id to get unique users per day
    daily_active_users = filtered_activity.groupby('activity_date')['user_id'].nunique().reset_index()
    
    # Rename the columns as per the expected output
    daily_active_users.columns = ['day', 'active_users']
    
    return daily_active_users
