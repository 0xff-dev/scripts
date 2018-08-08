
class Solve(object):

    def romanToInt(self, _str):
        """
        :type s: str
        :rtype: int
        """
        data_map = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
        }
        if len(_str) == 1:
            return data_map[_str]
        res = []
        flag = True
        for i in range(len(_str)-1):
            if (data_map[_str[i]] < data_map[_str[i+1]]):
                if flag:
                    res.append(-data_map[_str[i]])
                    flag = False
                else:
                    res.append(data_map[_str[i]])
                    flag = True
            elif data_map[_str[i]] == data_map[_str[i+1]]:
                if flag:
                    res.append(data_map[_str[i]])
                    flag = True
                else:
                    res.append(-data_map[_str[i]])
                    flag = False
            else:
                flag = True
                res.append(data_map[_str[i]])
        res.append(data_map[_str[-1]])
        return sum(res)

    def removeElement(self, nums, val):
        def inner_func():
            for i in nums:
                if i != val:
                    yield i
        length = len(list(inner_func()))
        nums += list(inner_func())
        while len(nums) > length:
            del nums[0]
        return len(nums)

    def rotate(self, matrix):
        return list(map(list, zip(*matrix[-1::-1])))

solve = Solve()
print(solve.romanToInt('CCCII'))
print(solve.romanToInt('XXIC'))
