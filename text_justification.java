import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<String> fullJustify(String[] words, int maxWidth) {
        List<String> result = new ArrayList<>();
        int index = 0;

        while (index < words.length) {
            int totalChars = words[index].length();
            int last = index + 1;
            while (last < words.length) {
                if (totalChars + words[last].length() + 1 > maxWidth) break;
                totalChars += words[last].length() + 1;
                last++;
            }

            StringBuilder builder = new StringBuilder();
            int gaps = last - index - 1;
            // If last line or number of gaps is zero (only one word fits in the line)
            if (last == words.length || gaps == 0) {
                for (int i = index; i < last; i++) {
                    builder.append(words[i]);
                    builder.append(' ');
                }
                builder.deleteCharAt(builder.length() - 1); // Remove the trailing space
                while (builder.length() < maxWidth) {
                    builder.append(' ');
                }
            } else {
                // Middle justified
                int spaces = (maxWidth - totalChars) / gaps;
                int extraSpaces = (maxWidth - totalChars) % gaps;
                for (int i = index; i < last; i++) {
                    builder.append(words[i]);
                    if (i < last - 1) {
                        for (int j = 0; j <= (spaces + ((i - index) < extraSpaces ? 1 : 0)); j++) {
                            builder.append(' ');
                        }
                    }
                }
            }
            result.add(builder.toString());
            index = last;
        }
        
        return result;
    }
}
