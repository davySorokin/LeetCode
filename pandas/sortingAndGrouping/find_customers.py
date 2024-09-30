import pandas as pd

def find_customers(customer: pd.DataFrame, product: pd.DataFrame) -> pd.DataFrame:
    # Get the unique list of products in the product table
    all_products = set(product['product_key'])
    
    # Group by customer_id and aggregate the products they bought
    customer_products = customer.groupby('customer_id')['product_key'].apply(set).reset_index()
    
    # Filter customers who bought all the products
    customers_who_bought_all = customer_products[customer_products['product_key'] == all_products]
    
    # Return the customer_id in the expected format
    return customers_who_bought_all[['customer_id']]
