import pandas as pd

def get_average_time(activity: pd.DataFrame) -> pd.DataFrame:
    # Separate the start and end activity types into two DataFrames
    start_times = activity[activity['activity_type'] == 'start'][['machine_id', 'process_id', 'timestamp']]
    end_times = activity[activity['activity_type'] == 'end'][['machine_id', 'process_id', 'timestamp']]

    # Merge start and end DataFrames on machine_id and process_id
    merged = pd.merge(start_times, end_times, on=['machine_id', 'process_id'], suffixes=('_start', '_end'))

    # Calculate the processing time as the difference between end and start timestamps
    merged['processing_time'] = merged['timestamp_end'] - merged['timestamp_start']

    # Group by machine_id and calculate the average processing time, rounding to 3 decimal places
    result = merged.groupby('machine_id')['processing_time'].mean().reset_index()
    result['processing_time'] = result['processing_time'].round(3)

    return result
