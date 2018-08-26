#!/usr/bin/env python
# coding=utf-8

import sys
import os

pwd = os.path.dirname(os.path.relpath(__file__))
sys.path.append(pwd+'../')
os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'project.settings')


import django
django.setup()


# 接下来就可以使用django的环境进行脚本的编写
