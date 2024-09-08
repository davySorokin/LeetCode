func climbStairs(n int) int {
    // If there are 1 or 2 steps, the number of ways to climb
    // them is equal to the number of steps (base cases).
    if n <= 2 {
        return n
    }

    // We use dynamic programming to solve the problem efficiently.
    // We are interested in the number of distinct ways to reach each step.
    // To reach step n, we could have come from step n-1 (1 step) or step n-2 (2 steps).

    // prev2 represents the number of ways to reach (n-2)th step.
    // prev1 represents the number of ways to reach (n-1)th step.
    prev1, prev2 := 2, 1

    // The idea is that at each step, we update the number of ways to reach
    // the current step (which is the sum of the previous two steps).
    for i := 3; i <= n; i++ {
        curr := prev1 + prev2 // The current step is reached by either from the previous step or two steps back.
        prev2 = prev1         // Move the previous two steps up.
        prev1 = curr          // Now, prev1 becomes the current step's number of ways.
    }

    // After the loop, prev1 will hold the number of ways to reach the nth step.
    return prev1
}
