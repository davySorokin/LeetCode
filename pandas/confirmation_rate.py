import pandas as pd

def confirmation_rate(signups: pd.DataFrame, confirmations: pd.DataFrame) -> pd.DataFrame:
    # Count total confirmation requests and confirmed actions
    total_requests = confirmations.groupby('user_id').size().reset_index(name='total_requests')
    confirmed_requests = confirmations[confirmations['action'] == 'confirmed'].groupby('user_id').size().reset_index(name='confirmed_requests')
    
    # Merge the two results to get both total requests and confirmed requests for each user
    merged = pd.merge(total_requests, confirmed_requests, on='user_id', how='left')
    
    # Fill NaN for users who have 0 confirmed requests
    merged['confirmed_requests'] = merged['confirmed_requests'].fillna(0)
    
    # Calculate confirmation rate
    merged['confirmation_rate'] = (merged['confirmed_requests'] / merged['total_requests']).round(2)
    
    # Merge with the signups table to include all users (users with no confirmation requests should have rate 0)
    result = pd.merge(signups[['user_id']], merged[['user_id', 'confirmation_rate']], on='user_id', how='left')
    
    # Fill NaN for users who have no confirmation requests
    result['confirmation_rate'] = result['confirmation_rate'].fillna(0).round(2)
    
    return result[['user_id', 'confirmation_rate']]

# Example usage (assuming signups and confirmations are DataFrames created elsewhere):
# signups = pd.DataFrame(...)
# confirmations = pd.DataFrame(...)
# output = confirmation_rate(signups, confirmations)
# print(output)
