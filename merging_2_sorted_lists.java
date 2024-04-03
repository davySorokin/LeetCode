class Solution {
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        // Create a dummy node which will be the start of the new list
        ListNode dummy = new ListNode(-1);
        
        // Create a pointer for the current position in the new list
        ListNode current = dummy;
        
        // While neither list is at the end, insert the smallest value node to the new list
        while (list1 != null && list2 != null) {
            if (list1.val < list2.val) {
                current.next = list1;
                list1 = list1.next;
            } else {
                current.next = list2;
                list2 = list2.next;
            }
            current = current.next;
        }
        
        // At least one of the lists is now null.
        // Add the remaining elements of the non-null list to the new list
        if (list1 != null) {
            current.next = list1;
        } else {
            current.next = list2;
        }
        
        // The first node is a dummy node, so return the next node
        return dummy.next;
    }
}
