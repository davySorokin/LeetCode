class Solution {
    func calPoints(_ operations: [String]) -> Int {
        var stack = [Int]()

        for op in operations {
            switch op {
            case "+":
                let last = stack[stack.count - 1]
                let secondLast = stack[stack.count - 2]
                stack.append(last + secondLast)
            case "D":
                stack.append(stack.last! * 2)
            case "C":
                stack.removeLast()
            default:
                stack.append(Int(op)!)
            }
        }

        return stack.reduce(0, +)
    }
}
