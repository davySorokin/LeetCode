/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function(head) {
    let current = head;

    // Traverse the linked list
    while (current && current.next) {
        // If current node's value equals the next node's value
        if (current.val === current.next.val) {
            // Skip the next node
            current.next = current.next.next;
        } else {
            // Otherwise, move to the next node
            current = current.next;
        }
    }

    return head;
};
