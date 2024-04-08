class Solution {
    public int myAtoi(String s) {
        if (s == null || s.length() == 0) return 0;

        int index = 0, total = 0, sign = 1;
        // 1. Remove Spaces
        while (index < s.length() && s.charAt(index) == ' ')
            index++;

        // 2. Handle signs
        if (index < s.length() && (s.charAt(index) == '+' || s.charAt(index) == '-'))
            sign = s.charAt(index++) == '+' ? 1 : -1;
        
        // 3. Convert number and avoid overflow
        while (index < s.length()) {
            int digit = s.charAt(index) - '0';
            if (digit < 0 || digit > 9) break;

            // Check if total will be overflow after 10 times and add digit
            if (Integer.MAX_VALUE / 10 < total || 
                Integer.MAX_VALUE / 10 == total && Integer.MAX_VALUE % 10 < digit)
                return sign == 1 ? Integer.MAX_VALUE : Integer.MIN_VALUE;

            total = 10 * total + digit;
            index++;
        }
        return total * sign;
    }
}
