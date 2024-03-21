class Solution {
    public boolean isPalindrome(int x) {
        // Negative numbers cannot be palindromes because of the '-' sign
        // Also, if the last digit is 0, the first digit must also be 0, only number 0 satisfy this condition
        if (x < 0 || (x % 10 == 0 && x != 0)) {
            return false;
        }

        int revertedNumber = 0;
        while (x > revertedNumber) {
            revertedNumber = revertedNumber * 10 + x % 10;
            x /= 10;
        }

        // The number is a palindrome if it's the same as its reverse
        // or if it's the same when excluding the middle digit (for odd length numbers)
        return x == revertedNumber || x == revertedNumber / 10;
    }
}
