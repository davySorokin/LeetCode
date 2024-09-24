import pandas as pd

def students_and_examinations(students: pd.DataFrame, subjects: pd.DataFrame, examinations: pd.DataFrame) -> pd.DataFrame:
    # Count the number of times each student attended an exam for a subject
    examination_count = examinations.groupby(['student_id', 'subject_name']).size().reset_index(name='attended_exams')
    
    # Cross join students and subjects to get all combinations of students and subjects
    student_subject = pd.merge(students, subjects, how='cross')
    
    # Left join with the examination_count to get the attendance data
    result_df = pd.merge(student_subject, examination_count, on=['student_id', 'subject_name'], how='left')
    
    # Fill NaN values in 'attended_exams' with 0 (for students who didn't attend any exam for a subject)
    result_df['attended_exams'] = result_df['attended_exams'].fillna(0).astype(int)
    
    # Sort the result by 'student_id' and 'subject_name'
    result_df = result_df[['student_id', 'student_name', 'subject_name', 'attended_exams']].sort_values(by=['student_id', 'subject_name'])
    
    return result_df
