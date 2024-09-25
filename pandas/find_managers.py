import pandas as pd

def find_managers(employee: pd.DataFrame) -> pd.DataFrame:
    # Group by managerId and count the number of employees under each manager
    manager_counts = employee.groupby('managerId').size().reset_index(name='count')
    
    # Filter managers who have at least 5 direct reports
    managers_with_five_or_more_reports = manager_counts[manager_counts['count'] >= 5]
    
    # Join back to the employee table to get the manager names
    result = pd.merge(managers_with_five_or_more_reports, employee, left_on='managerId', right_on='id')
    
    # Select only the manager names
    return result[['name']]
