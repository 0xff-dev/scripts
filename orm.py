#!/usr/bin/env python
# coding=utf-8
"""
python ORM使用, metaclass 隐式继承到子类
"""

class Field(object):

    def __init__(self, name, column_type):
        self.name = name
        self.column_type = column_type

    def __str__(self):
        return '<{}: {}>'.format(self.name, self.column_type)


class InterField(Field):

    def __init__(self, name):
        super(InterField, self).__init__(name, 'bigint')


class StringField(Field):

    def __init__(self, name):
        super(StringField, self).__init__(name, 'varchar(100)')


class ModelMetaclass(type):

    def __new__(cls, name, bases, attrs: dict):
        # Modal基类不管，直接处理子类
        if name == 'Model':
            return type.__new__(cls, name, bases, attrs)
        print('FOUND model: {}'.format(name))
        mapping = {}
        for key, value in attrs.items():
            if isinstance(value, Field):
                print('field {}===>{}'.format(key, value))
                mapping[key] = value
        for key in mapping.keys():
            attrs.pop(key)
        attrs['__mappings__'] = mapping
        attrs['__table__'] = name
        return type.__new__(cls, name, bases, attrs)

    
class Model(dict, metaclass=ModelMetaclass):

    def __init__(self, **kwargs):
        super(Model, self).__init__(**kwargs)

    def __getattr__(self, key):
        try:
            return self[key]
        except KeyError:
            raise AttributeError('Key<{}> not exists'.format(key))
    
    def __setattr__(self, key, value):
        self[key] = value

    def save(self):
        fields = []
        values = []
        args = []
        """
        id = IntegerField('id')
        """
        for key, value in self.__mappings__.items():
            print(key, ": ", value)
            fields.append(value.name)
            values.append('%s')
            args.append(getattr(self, key, None))
        sql = 'insert into {table}('.format(table=self.__table__)+\
              ','.join(fields)+") values("+\
              ','.join(values)+')'
        print('SQL: ', sql % tuple(args))


class User(Model):
    id = InterField('id')
    name = StringField('username')
    pwd = StringField('password')


user = User(id='123', name='Coco', pwd="asdfgh")
user.save()