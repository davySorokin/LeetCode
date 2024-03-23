class Solution {
    public int romanToInt(String s) {
        // Define a map for roman numerals
        Map<Character, Integer> romanValues = new HashMap<>();
        romanValues.put('I', 1);
        romanValues.put('V', 5);
        romanValues.put('X', 10);
        romanValues.put('L', 50);
        romanValues.put('C', 100);
        romanValues.put('D', 500);
        romanValues.put('M', 1000);
        
        int sum = 0;
        int prevValue = 0;
        
        // Start from the end of the string
        for (int i = s.length() - 1; i >= 0; i--) {
            int value = romanValues.get(s.charAt(i));
            
            // If the previous value is greater than the current value, we need to subtract
            if (value < prevValue) {
                sum -= value;
            } else {
                // Otherwise, we add the value
                sum += value;
            }
            
            // Update the previous value
            prevValue = value;
        }
        
        return sum;
    }
}
