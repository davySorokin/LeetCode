import pandas as pd

def sales_analysis(sales: pd.DataFrame, product: pd.DataFrame) -> pd.DataFrame:
    # Step 1: Find the first year of sale for each product (similar to the "temp" CTE)
    temp = sales.groupby('product_id')['year'].min().reset_index()
    temp = temp.rename(columns={'year': 'first_year'})
    
    # Step 2: Merge the Sales data with the "temp" table on product_id and year
    result = pd.merge(sales, temp, left_on=['product_id', 'year'], right_on=['product_id', 'first_year'])
    
    # Step 3: Select the required columns (product_id, first_year, quantity, price)
    result = result[['product_id', 'first_year', 'quantity', 'price']]
    
    # Return the result sorted by product_id (for consistency)
    return result.sort_values(by='product_id').reset_index(drop=True)
