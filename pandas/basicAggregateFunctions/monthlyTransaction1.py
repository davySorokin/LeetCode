import pandas as pd

def monthly_transactions(transactions: pd.DataFrame) -> pd.DataFrame:
    # Extract the month (YYYY-MM) from the trans_date
    transactions['month'] = transactions['trans_date'].astype(str).str[:7]
    
    # Calculate the total number of transactions and their total amount, including null countries
    trans_stats = transactions.groupby(['month', 'country'], dropna=False).agg(
        trans_count=('id', 'count'),
        trans_total_amount=('amount', 'sum')
    ).reset_index()
    
    # Calculate the number of approved transactions and their total amount
    approved_stats = transactions[transactions['state'] == 'approved'].groupby(['month', 'country'], dropna=False).agg(
        approved_count=('id', 'count'),
        approved_total_amount=('amount', 'sum')
    ).reset_index()

    # Merge the two DataFrames on month and country
    result = pd.merge(trans_stats, approved_stats, on=['month', 'country'], how='left')
    
    # Fill NaN values for the approved_count and approved_total_amount with 0
    result[['approved_count', 'approved_total_amount']] = result[['approved_count', 'approved_total_amount']].fillna(0)
    
    # Ensure approved_count is an integer
    result['approved_count'] = result['approved_count'].astype(int)

    return result
