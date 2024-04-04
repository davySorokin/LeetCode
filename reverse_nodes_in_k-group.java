class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode dummy = new ListNode(-1), pre = dummy, cur = head;
        dummy.next = head;

        while (cur != null) {
            ListNode start = cur, end = null;
            int count = 0;
            while (cur != null && count < k) {
                end = cur;
                cur = cur.next;
                count++;
            }
            if (count == k) {
                end.next = null;
                reverse(start);
                pre.next = end;
                start.next = cur;
                pre = start;
            } else {
                pre.next = start;
            }
        }
        return dummy.next;
    }

    private void reverse(ListNode head) {
        ListNode prev = null, cur = head;
        while (cur != null) {
            ListNode next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }
    }
}
