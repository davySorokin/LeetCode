import java.util.PriorityQueue;
import java.util.Comparator;

class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0) {
            return null;
        }
        
        PriorityQueue<ListNode> queue = new PriorityQueue<>(
            Comparator.comparingInt(n -> n.val)
        );

        // Add the first node of each list to the priority queue
        for (ListNode node : lists) {
            if (node != null) {
                queue.add(node);
            }
        }

        // Dummy head can help simplify edge cases when constructing the result list
        ListNode dummy = new ListNode(0);
        ListNode tail = dummy;

        while (!queue.isEmpty()) {
            ListNode minNode = queue.poll();
            tail.next = minNode;
            tail = tail.next;

            if (minNode.next != null) {
                queue.add(minNode.next);
            }
        }

        return dummy.next;
    }
}
