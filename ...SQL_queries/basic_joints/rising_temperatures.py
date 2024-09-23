import pandas as pd

def rising_temperature(weather: pd.DataFrame) -> pd.DataFrame:
    # Sort the DataFrame by the recordDate to ensure correct chronological order
    weather = weather.sort_values(by="recordDate")

    # Use the shift function to compare the temperature with the previous day
    weather['prev_temperature'] = weather['temperature'].shift(1)
    
    # Filter the rows where the temperature is greater than the previous day's temperature
    result = weather[weather['temperature'] > weather['prev_temperature']]

    # Return only the 'id' column as per the required output
    return result[['id']]
