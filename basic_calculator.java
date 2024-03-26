class Solution {
    public int calculate(String s) {
        Stack<Integer> stack = new Stack<>();
        int result = 0; // Current result
        int number = 0; // Current number
        int sign = 1; // 1 means positive, -1 means negative

        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (Character.isDigit(c)) {
                number = 10 * number + (c - '0');
            } else if (c == '+') {
                result += sign * number;
                number = 0;
                sign = 1;
            } else if (c == '-') {
                result += sign * number;
                number = 0;
                sign = -1;
            } else if (c == '(') {
                // Push the result and the sign onto the stack, for later
                // We push the result first, then the sign
                stack.push(result);
                stack.push(sign);
                // Reset the sign and result for the value in the parenthesis
                sign = 1;
                result = 0;
            } else if (c == ')') {
                // End of parenthesis, process the current number and add to the result before the parenthesis
                result += sign * number;
                number = 0;
                result *= stack.pop(); // stack.pop() is the sign before the parenthesis
                result += stack.pop(); // stack.pop() now contains the result before the parenthesis
            }
        }
        // Add the last number
        if (number != 0) result += sign * number;

        return result;
    }
}
