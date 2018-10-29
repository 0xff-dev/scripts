from functools import wraps

"""
python 单例模式设计
"""

# 1. __new__

class SingletonOne(object):

    def __new__(cls, **kwargs):
        if not hasattr(cls, '_instance'):
            cls._instace = super(SingletonOne, cls).__new__(cls)
        return cls._instace

    def __init__(self, **kwargs):
        self.name = kwargs['name']
        self.age = kwargs['age']

    def __str__(self):
        return self.name+str(self.age)

test_one = SingletonOne(name='Cocod', age=12)
print(test_one)


# 2. decorator

def SingletonTwo(cls):
    _instance = {}
    @wraps(cls)
    def _singleton(*args, **kwargs):
        if cls not in _instance:
            _instance[cls] = cls(*args, **kwargs)
        return _instance[cls]
    return _singleton


@SingletonTwo
class Person(object):
    name = 'coco'


obj1 = Person()
obj2 = Person()
print(obj1 is obj2)   # true


# 3. metaclass, 元类需要多学习, 明天继续