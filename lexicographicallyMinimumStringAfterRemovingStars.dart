class Solution {
  String clearStars(String s) {
    int n = s.length;
    List<List<int>> g = List.generate(26, (_) => []);
    List<bool> deleted = List.filled(n, false);
    List<String> chars = s.split('');

    for (int i = 0; i < n; i++) {
      if (chars[i] == '*') {
        deleted[i] = true;

        for (int j = 0; j < 26; j++) {
          if (g[j].isNotEmpty) {
            int idx = g[j].removeLast();
            deleted[idx] = true;
            break;
          }
        }
      } else {
        int chIdx = chars[i].codeUnitAt(0) - 'a'.codeUnitAt(0);
        g[chIdx].add(i);
      }
    }

    StringBuffer result = StringBuffer();
    for (int i = 0; i < n; i++) {
      if (!deleted[i] && chars[i] != '*') {
        result.write(chars[i]);
      }
    }

    return result.toString();
  }
}
