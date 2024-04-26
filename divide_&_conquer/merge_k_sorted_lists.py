import heapq

class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        # Initialize the heap with tuples of the form (node value, index, node)
        heap = [(l.val, i, l) for i, l in enumerate(lists) if l]
        heapq.heapify(heap)
        
        # Create a dummy node to start the merged list
        dummy = ListNode(0)
        current = dummy
        
        # Merge the k lists
        while heap:
            # Pop the smallest element from the heap (by node value)
            _, i, node = heapq.heappop(heap)
            current.next = node
            current = current.next
            # If there is a next node, push it to the heap with updated index
            if node.next:
                heapq.heappush(heap, (node.next.val, i, node.next))
        
        return dummy.next
