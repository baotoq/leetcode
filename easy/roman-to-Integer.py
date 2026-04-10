class Solution:
    def romanToInt(self, str: str) -> int:
        result = 0
        dic = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
        for i in range(len(str)):
            if i < len(str) - 1 and dic[str[i]] < dic[str[i + 1]]:
                result -= dic[str[i]]
            else:
                result += dic[str[i]]
        return result
