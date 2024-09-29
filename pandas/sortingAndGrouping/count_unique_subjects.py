import pandas as pd

def count_unique_subjects(teacher: pd.DataFrame) -> pd.DataFrame:
    # Group by teacher_id and subject_id, remove duplicate subjects taught by the same teacher in different departments
    result = teacher.groupby('teacher_id')['subject_id'].nunique().reset_index()
    
    # Rename columns as per the expected output
    result.columns = ['teacher_id', 'cnt']
    
    return result
