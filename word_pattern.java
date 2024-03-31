import java.util.HashMap;
import java.util.Map;

class Solution {
    public boolean wordPattern(String pattern, String s) {
        String[] words = s.split(" "); // Split the string into words
        if (words.length != pattern.length()) return false; // Early check

        Map<Character, String> charToWord = new HashMap<>();
        Map<String, Character> wordToChar = new HashMap<>();

        for (int i = 0; i < pattern.length(); i++) {
            char c = pattern.charAt(i);
            String word = words[i];

            // Check if the mapping from char to word already exists
            if (charToWord.containsKey(c)) {
                if (!charToWord.get(c).equals(word)) return false; // Not following the pattern
            } else {
                charToWord.put(c, word);
            }

            // Check if the mapping from word to char already exists
            if (wordToChar.containsKey(word)) {
                if (wordToChar.get(word) != c) return false; // Not following the pattern
            } else {
                wordToChar.put(word, c);
            }
        }

        return true;
    }
}
