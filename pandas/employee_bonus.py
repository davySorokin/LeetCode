import pandas as pd

def employee_bonus(employee: pd.DataFrame, bonus: pd.DataFrame) -> pd.DataFrame:
    # Perform an outer join to keep all employees, even those without bonuses
    result = pd.merge(employee, bonus, on='empId', how='left')
    
    # Fill NaN values in the bonus column with pd.NA (to represent nulls)
    result['bonus'] = result['bonus'].fillna(pd.NA)
    
    # Filter rows where the bonus is either null or less than 1000
    result = result[(result['bonus'].isna()) | (result['bonus'] < 1000)]
    
    # Select only the name and bonus columns
    result = result[['name', 'bonus']]
    
    return result
