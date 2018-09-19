# coding=utf-8

import json
import xlrd


sheet_name = 'Sheet1'


def read_file(file_path):
    excel = xlrd.open_workbook(file_path)
    sheet1 = excel.sheet_by_name('Sheet1')
    cols = sheet1.ncols
    names = []
    res = []
    # 找好name
    for i in range(cols):
        names.append(set(sheet1.col_values(i)))
    if cols > 1:
        # 第一层的 [{}, {}, {}], 接下来的就把children的[]拿出来即可
        for name in names[0]:
            res.append({"name": name, "children": []})
        
        childrens = dict(map(lambda x: (x["name"], x["children"]), res))
        def insert_children(zip_datas: list, last_cols: bool):
            nonlocal childrens
            if last_cols:
                for data in zip_datas:
                    childrens[data[0]].append({"name": data[1]})
            else:
                for data in zip_datas:
                    childrens[data[0]].append({"name": data[1], "children": []})
                t_childrens = {}
                for data in childrens.keys():
                    for data_list in childrens[data]:
                        t_childrens[data_list["name"]] = data_list["children"]
                childrens = t_childrens
        for i in range(cols-1):
            zip_datas = list(set(zip(sheet1.col_values(i), sheet1.col_values(i+1))))
            insert_children(zip_datas, True if i == cols-2 else False)
    else:
         res = list(map(lambda x: {"name": x}, names[0]))
    print(json.dumps(res, ensure_ascii=False))
read_file('data.xlsx')
