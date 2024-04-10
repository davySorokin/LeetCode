import java.util.Deque;
import java.util.LinkedList;
import java.util.Arrays;

class Solution {
    public int[] deckRevealedIncreasing(int[] deck) {
        int n = deck.length;
        Deque<Integer> indexQueue = new LinkedList<>();
        for (int i = 0; i < n; i++) {
            indexQueue.addLast(i);
        }
        
        Arrays.sort(deck);
        int[] result = new int[n];
        
        for (int card : deck) {
            result[indexQueue.pollFirst()] = card; // Place the card at the next available position
            if (!indexQueue.isEmpty()) {
                indexQueue.addLast(indexQueue.pollFirst()); // Move the next index to the end of the queue
            }
        }
        
        return result;
    }
}
