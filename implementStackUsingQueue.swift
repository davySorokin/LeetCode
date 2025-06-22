/**
 * Your MyStack object will be instantiated and called as such:
 * let obj = MyStack()
 * obj.push(x)
 * let ret_2: Int = obj.pop()
 * let ret_3: Int = obj.top()
 * let ret_4: Bool = obj.empty()
 */

 class MyStack {
    private var q1 = [Int]()
    private var q2 = [Int]()

    init() {}

    func push(_ x: Int) {
        q2.append(x)
        while !q1.isEmpty {
            q2.append(q1.removeFirst())
        }
        swap(&q1, &q2)
    }

    func pop() -> Int {
        return q1.removeFirst()
    }

    func top() -> Int {
        return q1.first!
    }

    func empty() -> Bool {
        return q1.isEmpty
    }
}
