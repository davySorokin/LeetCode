/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function(headA, headB) {
    if (!headA || !headB) return null;

    let pointerA = headA;
    let pointerB = headB;

    // Traverse both lists. Once one pointer reaches the end, redirect it to the other list's head.
    // If they intersect, they will eventually meet at the intersection node.
    // If they don't intersect, both will reach null after traversing both lists.
    while (pointerA !== pointerB) {
        pointerA = pointerA ? pointerA.next : headB;
        pointerB = pointerB ? pointerB.next : headA;
    }

    // Either they meet at the intersection node or both become null (no intersection).
    return pointerA;
};
