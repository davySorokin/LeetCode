class Solution {
    public String removeKdigits(String num, int k) {
        if (k == num.length()) {
            return "0";
        }
        
        StringBuilder sb = new StringBuilder();
        
        for (char c : num.toCharArray()) {
            // While stack is not empty and the last digit in stack is greater than the current digit
            // and we still have to remove digits
            while (sb.length() > 0 && k > 0 && sb.charAt(sb.length() - 1) > c) {
                sb.deleteCharAt(sb.length() - 1);
                k--;
            }
            sb.append(c);
        }
        
        // Remove any remaining digits from the end
        while (k > 0) {
            sb.deleteCharAt(sb.length() - 1);
            k--;
        }
        
        // Remove leading zeros
        while (sb.length() > 1 && sb.charAt(0) == '0') {
            sb.deleteCharAt(0);
        }
        
        return sb.toString();
    }
}
