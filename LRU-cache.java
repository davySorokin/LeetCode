import java.util.HashMap;
import java.util.Map;

class LRUCache {

    private Map<Integer, ListNode> map;
    private DoubleLinkedList cache;
    private int capacity;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        map = new HashMap<>();
        cache = new DoubleLinkedList();
    }
    
    public int get(int key) {
        if (!map.containsKey(key)) {
            return -1;
        }
        // Move the accessed node to the head (mark as most recently used)
        ListNode node = map.get(key);
        cache.moveToHead(node);
        return node.value;
    }
    
    public void put(int key, int value) {
        ListNode node = map.get(key);
        if (node == null) {
            // If the key doesn't exist, create a new ListNode
            ListNode newNode = new ListNode(key, value);
            // Add to the HashMap and the actual list that represents the cache
            map.put(key, newNode);
            cache.addToHead(newNode);
            if (map.size() > capacity) {
                // If the cache exceeded the capacity, remove the LRU entry from map and list
                ListNode last = cache.removeLast();
                map.remove(last.key);
            }
        } else {
            // If the key exists, update the value and move it to the head
            node.value = value;
            cache.moveToHead(node);
        }
    }

    // Node of a doubly linked list
    class ListNode {
        int key;
        int value;
        ListNode prev;
        ListNode next;

        public ListNode() {}

        public ListNode(int key, int value) {
            this.key = key;
            this.value = value;
        }
    }

    // A doubly linked list to store the order of the keys
    class DoubleLinkedList {
        private ListNode head;
        private ListNode tail;

        public DoubleLinkedList() {
            head = new ListNode();
            tail = new ListNode();
            head.next = tail;
            tail.prev = head;
        }

        public void addToHead(ListNode node) {
            node.prev = head;
            node.next = head.next;
            head.next.prev = node;
            head.next = node;
        }

        public void removeNode(ListNode node) {
            node.prev.next = node.next;
            node.next.prev = node.prev;
        }

        public void moveToHead(ListNode node) {
            removeNode(node);
            addToHead(node);
        }

        public ListNode removeLast() {
            if (tail.prev == head) {
                return null;
            }
            ListNode last = tail.prev;
            removeNode(last);
            return last;
        }
    }
}
