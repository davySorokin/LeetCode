#include <string>
#include <unordered_map>
using namespace std;

class Solution {
public:
    bool isIsomorphic(string s, string t) {
        unordered_map<char, char> sTot;
        unordered_map<char, char> tTos;
        
        for (int i = 0; i < s.length(); ++i) {
            char sChar = s[i];
            char tChar = t[i];
            
            // Check if there is a conflicting mapping from s to t or from t to s
            if ((sTot.count(sChar) && sTot[sChar] != tChar) || 
                (tTos.count(tChar) && tTos[tChar] != sChar)) {
                return false;
            }
            
            sTot[sChar] = tChar;
            tTos[tChar] = sChar;
        }
        
        return true;
    }
};
