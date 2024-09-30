import pandas as pd

def find_classes(courses: pd.DataFrame) -> pd.DataFrame:
    # Group by class and count the number of students in each class
    class_counts = courses.groupby('class').size().reset_index(name='student_count')
    
    # Filter the classes with at least 5 students
    result = class_counts[class_counts['student_count'] >= 5]
    
    # Return only the 'class' column as per the expected output format
    return result[['class']]
