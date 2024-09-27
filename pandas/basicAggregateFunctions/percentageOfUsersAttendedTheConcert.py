import pandas as pd

def users_percentage(users: pd.DataFrame, register: pd.DataFrame) -> pd.DataFrame:
    # Counting total users
    total_users = users['user_id'].nunique()
    
    # Counting the number of users per contest
    contest_counts = register.groupby('contest_id').agg(registered_users=('user_id', 'nunique')).reset_index()
    
    # Calculating the percentage of users per contest
    contest_counts['percentage'] = (contest_counts['registered_users'] / total_users) * 100
    
    # Rounding the percentage to two decimal places
    contest_counts['percentage'] = contest_counts['percentage'].round(2)
    
    # Sorting by percentage in descending order and by contest_id in ascending order in case of a tie
    contest_counts = contest_counts.sort_values(by=['percentage', 'contest_id'], ascending=[False, True]).reset_index(drop=True)
    
    return contest_counts[['contest_id', 'percentage']]
