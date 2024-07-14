package main

import (
    "fmt"
)

func maxCompatibilitySum(students [][]int, mentors [][]int) int {
    m := len(students)
    n := len(students[0])
    permutations := permute(generateIndices(m))
    maxScore := 0

    for _, perm := range permutations {
        currentScore := 0
        for i, studentIdx := range perm {
            currentScore += compatibilityScore(students[studentIdx], mentors[i], n)
        }
        if currentScore > maxScore {
            maxScore = currentScore
        }
    }
    return maxScore
}

func generateIndices(m int) []int {
    indices := make([]int, m)
    for i := 0; i < m; i++ {
        indices[i] = i
    }
    return indices
}

func permute(nums []int) [][]int {
    var result [][]int
    var backtrack func(int)
    backtrack = func(start int) {
        if start == len(nums) {
            perm := make([]int, len(nums))
            copy(perm, nums)
            result = append(result, perm)
        }
        for i := start; i < len(nums); i++ {
            nums[start], nums[i] = nums[i], nums[start]
            backtrack(start + 1)
            nums[start], nums[i] = nums[i], nums[start]
        }
    }
    backtrack(0)
    return result
}

func compatibilityScore(student, mentor []int, n int) int {
    score := 0
    for i := 0; i < n; i++ {
        if student[i] == mentor[i] {
            score++
        }
    }
    return score
}
