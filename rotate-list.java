class Solution {
    public ListNode rotateRight(ListNode head, int k) {
        if (head == null || head.next == null || k == 0) {
            return head;
        }

        // Find the length of the list and make it circular
        int length = 1; // Starts from 1 to include the head node
        ListNode oldTail = head;
        while (oldTail.next != null) {
            oldTail = oldTail.next;
            length++;
        }
        oldTail.next = head; // Makes the list circular

        // Find the new tail: (length - k % length) steps away from the head
        ListNode newTail = head;
        for (int i = 0; i < length - k % length - 1; i++) {
            newTail = newTail.next;
        }

        // Set the new head and break the circle
        ListNode newHead = newTail.next;
        newTail.next = null;

        return newHead;
    }
}
