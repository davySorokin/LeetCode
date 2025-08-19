class Solution {
  List<String> findRepeatedDnaSequences(String s) {
    if (s.length < 10) return [];
    const enc = {'A': 0, 'C': 1, 'G': 2, 'T': 3};
    final seen = <int>{};
    final added = <int>{};
    final res = <String>[];
    final mask = (1 << 20) - 1;
    var x = 0;

    for (var i = 0; i < s.length; i++) {
      x = ((x << 2) & mask) | enc[s[i]]!;
      if (i >= 9) {
        if (seen.contains(x)) {
          if (!added.contains(x)) {
            res.add(s.substring(i - 9, i + 1));
            added.add(x);
          }
        } else {
          seen.add(x);
        }
      }
    }
    return res;
  }
}
