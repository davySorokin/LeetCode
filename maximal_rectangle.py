class Solution:
    def maximalRectangle(self, matrix):
        if not matrix or not matrix[0]:
            return 0

        max_area = 0
        height = [0] * (len(matrix[0]) + 1)

        for row in matrix:
            for i in range(len(matrix[0])):
                height[i] = height[i] + 1 if row[i] == '1' else 0

            stack = [-1]

            for i in range(len(height)):
                while height[i] < height[stack[-1]]:
                    h = height[stack.pop()]
                    w = i - stack[-1] - 1
                    max_area = max(max_area, h * w)
                stack.append(i)

        return max_area
