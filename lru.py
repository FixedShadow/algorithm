class listNode:
    def __init__(self, key=0, value=0, pre=None, next=None):
        self.key = key
        self.value = value
        self.pre = pre
        self.next = next


class LRUCache:
    def __init__(self, capacity: int):
        self.count = 0
        self.limit = capacity
        self.head = listNode()
        self.tail = listNode()
        self.head.next = self.tail
        self.tail.pre = self.head
        self.element = dict()

    def get(self, key: int) -> int:
        if key in self.element:
            self.move_to_front(key)
            return self.element[key].value
        return -1

    def put(self, key: int, value: int) -> None:
        if key in self.element:
            self.element[key].value = value
            self.move_to_front(key)
            return
        if self.count < self.limit:
            self.add_node(key, value)
            return
        self.remove_tail_node()
        self.add_node(key, value)

    def move_to_front(self, key: int):
        node = self.element[key]
        tmp = node.pre
        tmp.next = node.next
        node.next.pre = tmp
        node.next = self.head.next
        self.head.next.pre = node
        self.head.next = node
        node.pre = self.head

    def add_node(self, key: int, value: int):
        node = listNode(key, value)
        node.next = self.head.next
        self.head.next.pre = node
        self.head.next = node
        node.pre = self.head
        self.count += 1
        self.element[key] = node

    def remove_tail_node(self):
        node = self.tail.pre
        tmp = self.tail.pre.pre
        tmp.next = self.tail
        self.tail.pre = tmp
        self.element.pop(node.key)
        self.count -= 1

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)