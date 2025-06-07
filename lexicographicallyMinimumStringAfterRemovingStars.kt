class Solution {
    fun clearStars(s: String): String {
        val n = s.length
        val g = Array(26) { ArrayDeque<Int>() }
        val deleted = BooleanArray(n)

        for (i in s.indices) {
            if (s[i] == '*') {
                deleted[i] = true
                for (j in 0 until 26) {
                    if (g[j].isNotEmpty()) {
                        val idx = g[j].removeLast()
                        deleted[idx] = true
                        break
                    }
                }
            } else {
                g[s[i] - 'a'].add(i)
            }
        }

        val result = StringBuilder()
        for (i in s.indices) {
            if (!deleted[i] && s[i] != '*') {
                result.append(s[i])
            }
        }

        return result.toString()
    }
}
