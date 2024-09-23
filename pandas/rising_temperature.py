import pandas as pd

def rising_temperature(weather: pd.DataFrame) -> pd.DataFrame:
    # Sort the DataFrame by the recordDate to ensure correct chronological order
    weather = weather.sort_values(by="recordDate")

    # Use the shift function to get the previous recordDate and temperature
    weather['prev_recordDate'] = weather['recordDate'].shift(1)
    weather['prev_temperature'] = weather['temperature'].shift(1)
    
    # Convert recordDate to datetime to calculate the difference in days
    weather['recordDate'] = pd.to_datetime(weather['recordDate'])
    weather['prev_recordDate'] = pd.to_datetime(weather['prev_recordDate'])
    
    # Calculate the difference in days between the current and previous recordDate
    weather['date_diff'] = (weather['recordDate'] - weather['prev_recordDate']).dt.days

    # Filter for rows where the temperature is higher than the previous day and the dates are consecutive
    result = weather[(weather['temperature'] > weather['prev_temperature']) & (weather['date_diff'] == 1)]

    # Return only the 'id' column as per the required output
    return result[['id']]
