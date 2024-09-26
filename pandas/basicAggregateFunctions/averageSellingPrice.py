import pandas as pd

def average_selling_price(prices: pd.DataFrame, units_sold: pd.DataFrame) -> pd.DataFrame:
    # Merge the prices and units_sold DataFrames based on product_id and purchase_date falling within the price period
    merged_df = pd.merge(units_sold, prices, on='product_id', how='right')

    # Filter for rows where the purchase_date is between start_date and end_date
    valid_rows = merged_df[(merged_df['purchase_date'] >= merged_df['start_date']) & 
                           (merged_df['purchase_date'] <= merged_df['end_date'])]

    # Calculate the total revenue for each sale (units * price)
    valid_rows['total_revenue'] = valid_rows['units'] * valid_rows['price']
    
    # Group by product_id and calculate the total revenue and total units sold for each product
    total_revenue_per_product = valid_rows.groupby('product_id')['total_revenue'].sum().reset_index(name='total_revenue')
    total_units_per_product = valid_rows.groupby('product_id')['units'].sum().reset_index(name='total_units')

    # Merge the total revenue and total units into a single DataFrame
    result_df = pd.merge(total_revenue_per_product, total_units_per_product, on='product_id', how='outer')
    
    # Calculate the average selling price for each product, handling products with no sales
    result_df['average_price'] = (result_df['total_revenue'] / result_df['total_units']).fillna(0).round(2)
    
    # Get all unique product_ids from the prices table and merge with the result to handle products with no sales
    all_products = prices[['product_id']].drop_duplicates()
    final_result = pd.merge(all_products, result_df[['product_id', 'average_price']], on='product_id', how='left')

    # Fill any NaN average prices (for products with no sales) with 0
    final_result['average_price'] = final_result['average_price'].fillna(0).round(2)
    
    return final_result[['product_id', 'average_price']]
