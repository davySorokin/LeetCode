<?php
class Solution {
    /**
     * @param String $s
     * @return Boolean
     */
    function repeatedSubstringPattern($s) {
        $len = strlen($s);
        
        for ($i = 1; $i <= $len / 2; $i++) {
            if ($len % $i == 0) {
                $substring = substr($s, 0, $i);
                $repeated = str_repeat($substring, $len / $i);
                if ($repeated === $s) {
                    return true;
                }
            }
        }
        
        return false;
    }
}
?>
