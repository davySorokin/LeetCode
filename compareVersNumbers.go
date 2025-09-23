import (
    "strconv"
    "strings"
)

func compareVersion(version1 string, version2 string) int {
    v1Parts := strings.Split(version1, ".")
    v2Parts := strings.Split(version2, ".")
    maxLen := len(v1Parts)
    if len(v2Parts) > maxLen {
        maxLen = len(v2Parts)
    }

    for i := 0; i < maxLen; i++ {
        var num1, num2 int
        if i < len(v1Parts) {
            num1, _ = strconv.Atoi(v1Parts[i])
        }
        if i < len(v2Parts) {
            num2, _ = strconv.Atoi(v2Parts[i])
        }
        if num1 < num2 {
            return -1
        } else if num1 > num2 {
            return 1
        }
    }
    return 0
}
