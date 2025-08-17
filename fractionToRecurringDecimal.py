class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        if numerator == 0:
            return "0"

        # Sign
        neg = (numerator < 0) ^ (denominator < 0)
        n, d = abs(numerator), abs(denominator)

        # Integer part
        integer_part = n // d
        rem = n % d
        res = ["-" if neg else "", str(integer_part)]

        if rem == 0:
            return "".join(res)

        res.append(".")
        # remainder -> index in res after the decimal point begins
        seen = {}
        while rem != 0:
            if rem in seen:
                idx = seen[rem]
                res.insert(idx, "(")
                res.append(")")
                break
            seen[rem] = len(res)
            rem *= 10
            res.append(str(rem // d))
            rem %= d

        return "".join(res)
