import pandas as pd

def immediate_food_delivery(delivery: pd.DataFrame) -> pd.DataFrame:
    # Get the first order per customer by finding the minimum order_date
    first_orders = delivery.loc[delivery.groupby('customer_id')['order_date'].idxmin()]
    
    # Determine if the first order is immediate (order_date equals customer_pref_delivery_date)
    first_orders['is_immediate'] = (first_orders['order_date'] == first_orders['customer_pref_delivery_date']).astype(int)
    
    # Calculate the percentage of immediate first orders
    immediate_percentage = first_orders['is_immediate'].mean() * 100
    
    # Return the result as a DataFrame with the percentage rounded to 2 decimal places
    result = pd.DataFrame({'immediate_percentage': [round(immediate_percentage, 2)]})
    
    return result
