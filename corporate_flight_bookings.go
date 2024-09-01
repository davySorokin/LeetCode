func corpFlightBookings(bookings [][]int, n int) []int {
    answer := make([]int, n)

    for _, booking := range bookings {
        first := booking[0] - 1
        last := booking[1] - 1
        seats := booking[2]

        answer[first] += seats
        if last + 1 < n {
            answer[last + 1] -= seats
        }
    }

    for i := 1; i < n; i++ {
        answer[i] += answer[i - 1]
    }

    return answer
}
