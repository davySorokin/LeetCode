import pandas as pd

def project_employees_i(project: pd.DataFrame, employee: pd.DataFrame) -> pd.DataFrame:
    # Merging the project and employee dataframes on employee_id
    merged_df = pd.merge(project, employee, on='employee_id')
    
    # Grouping by project_id and calculating the average of experience_years
    result_df = merged_df.groupby('project_id').agg(average_years=('experience_years', 'mean')).reset_index()
    
    # Rounding the average experience to 2 decimal places
    result_df['average_years'] = result_df['average_years'].round(2)
    
    return result_df
