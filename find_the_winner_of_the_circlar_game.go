func findTheWinner(n int, k int) int {
    winner := 0  // Start with index 0, which corresponds to the first friend

    // Iterate to compute the winner using the Josephus problem solution
    for i := 2; i <= n; i++ {
        winner = (winner + k) % i
    }

    return winner + 1  // Convert 0-based index to 1-based index
}

/*
This solution is based on the Josephus problem, where the position of the winner can be computed using a mathematical approach.
The time complexity is O(n) as we iterate through all friends. The space complexity is O(1) because we only use a constant amount of extra space. This approach meets the linear time and constant space constraints.
*/
