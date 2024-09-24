import pandas as pd

def students_and_examinations(students: pd.DataFrame, subjects: pd.DataFrame, examinations: pd.DataFrame) -> pd.DataFrame:
    # Create the cross join between students and subjects
    cross_join = students.assign(key=1).merge(subjects.assign(key=1), on='key').drop('key', axis=1)
    
    # Left join the cross join result with the examinations table to get the attendance count
    result = cross_join.merge(examinations.groupby(['student_id', 'subject_name']).size().reset_index(name='attended_exams'),
                              on=['student_id', 'subject_name'], how='left').fillna(0)
    
    # Convert the attendance counts to integers
    result['attended_exams'] = result['attended_exams'].astype(int)
    
    # Sort by student_id and subject_name
    result = result.sort_values(by=['student_id', 'subject_name'])
    
    # Return the final result
    return result[['student_id', 'student_name', 'subject_name', 'attended_exams']]
