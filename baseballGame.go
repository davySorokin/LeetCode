func calPoints(operations []string) int {
    stack := []int{}

    for _, op := range operations {
        switch op {
        case "+":
            n := len(stack)
            stack = append(stack, stack[n-1]+stack[n-2])
        case "D":
            stack = append(stack, 2*stack[len(stack)-1])
        case "C":
            stack = stack[:len(stack)-1]
        default:
            // Convert string to integer
            val, _ := strconv.Atoi(op)
            stack = append(stack, val)
        }
    }

    total := 0
    for _, score := range stack {
        total += score
    }

    return total
}
