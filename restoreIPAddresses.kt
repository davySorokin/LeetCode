class Solution {
    fun restoreIpAddresses(s: String): List<String> {
        val result = mutableListOf<String>()
        if (s.length < 4 || s.length > 12) return result

        fun backtrack(start: Int, dots: Int, current: String) {
            // If we've placed 3 dots and we're at the end of the string, it's valid
            if (dots == 3) {
                val part = s.substring(start)
                if (isValid(part)) {
                    result.add("$current$part")
                }
                return
            }

            // Try adding the dot after every 1, 2, or 3 digit substring
            for (i in start + 1 until s.length + 1) {
                val part = s.substring(start, i)
                if (isValid(part)) {
                    backtrack(i, dots + 1, "$current$part.")
                }
            }
        }

        // Helper function to validate a part
        fun isValid(part: String): Boolean {
            if (part.length > 3) return false
            if (part.length > 1 && part[0] == '0') return false
            return part.toIntOrNull()?.let { it in 0..255 } ?: false
        }

        // Start backtracking
        backtrack(0, 0, "")
        return result
    }
}

/*
fun main() {
    val solution = Solution()
    println(solution.restoreIpAddresses("25525511135"))
    println(solution.restoreIpAddresses("0000"))
    println(solution.restoreIpAddresses("101023"))
}
*/
