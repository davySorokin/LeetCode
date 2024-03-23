class Solution {
    public void reorderList(ListNode head) {
        if (head == null || head.next == null) return;

        // Find the middle of the list
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null && fast.next.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }

        // Reverse the second half of the list
        ListNode prev = null, curr = slow.next, tmp;
        while (curr != null) {
            tmp = curr.next;
            curr.next = prev;
            prev = curr;
            curr = tmp;
        }
        slow.next = null; // Split the list into two parts

        // Merge the two halves
        ListNode first = head, second = prev;
        while (second != null) {
            tmp = first.next;
            first.next = second;
            first = second;
            second = tmp;
        }
    }
}
