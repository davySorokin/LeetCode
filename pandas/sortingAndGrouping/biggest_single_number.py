import pandas as pd

def biggest_single_number(my_numbers: pd.DataFrame) -> pd.DataFrame:
    # Group by 'num' and count the occurrences of each number
    num_counts = my_numbers.groupby('num').size().reset_index(name='count')
    
    # Filter numbers that appear only once
    single_numbers = num_counts[num_counts['count'] == 1]
    
    if not single_numbers.empty:
        # Find the largest single number
        largest_single = single_numbers['num'].max()
        return pd.DataFrame({'num': [largest_single]})
    else:
        # Return null if no single number exists
        return pd.DataFrame({'num': [None]})
