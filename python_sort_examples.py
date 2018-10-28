#!/usr/bin/env python
# coding=utf-8

from random import randint

"""
[].sort(), 对本身进行修改
sorted([]) 返回一个排序后的[], 本身不做修改

sorted([], key, reverse)
"""

# 1, 数字排序
data = [randint(1, 100) for index in range(15)]
print(sorted(data))

# 2. 字典的排序, 先根据name排序，name相同按年龄反序排序
data = [
    {"name": "coco", "age": 22},
    {"name": "asdf", "age": 21},
    {"name": "coco", "age": 20}
]
"""
Expectation: [{'name': 'asdf', 'age': 21}, 
              {'name': 'coco', 'age': 22},
              {'name': 'coco', 'age': 20}
            ] 
"""
print(sorted(data, key=lambda x: (x['name'], -x['age'])))


# 3. 字典的key分析排序
data = {
    "coco-123": "coco",
    "coco-012": "coco",
    "asd-123": "asddf"
}

"""
根据data的key按照-分割的排序
Expectation: [('asd-123', 'asddf'),
              ('coco-012', 'coco'),
              ('coco-123', 'coco')
            ]
            or
            {
              'asd-123': 'asddf',
              'coco-012': 'coco',
              'coco-123': 'coco'
            }

"""
print(dict(sorted(data.items(), key=lambda x: tuple(x[0].split('-')))))
