import json
import xlrd
from collections import OrderedDict

sheet_name = 'Sheet1'


def read_file(file_path, sheet_name=sheet_name):
    excel = xlrd.open_workbook(file_path)
    sheet1 = excel.sheet_by_name(sheet_name)
    cols = sheet1.ncols
    names = []
    res = []
    # 找好name
    for i in range(cols):
        tmp_data = sheet1.col_values(i)
        tmp = list(set(tmp_data))
        tmp.sort(key=tmp_data.index)
        names.append(tmp)
    if cols > 1:
        for name in names[0]:
            res.append({"name": name, "children": []})
        
        childrens = OrderedDict(map(lambda x: (x["name"], x["children"]), res))
        def insert_children(zip_datas: list, last_cols: bool):
            nonlocal childrens
            if last_cols:
                for data in zip_datas:
                    childrens[data[0]].append({"name": data[1]})
            else:
                for data in zip_datas:
                    childrens[data[0]].append({"name": data[1], "children": []})
                t_childrens = OrderedDict()
                for data in childrens.keys():
                    for data_list in childrens[data]:
                        t_childrens[data_list["name"]] = data_list["children"]
                childrens = t_childrens
        for i in range(cols-1):
            list_datas = list(zip(sheet1.col_values(i), sheet1.col_values(i+1)))
            zip_datas = list(set(list_datas))
            zip_datas.sort(key=list_datas.index)
            insert_children(zip_datas, True if i == cols-2 else False)
    else:
        res = list(map(lambda x: {"name": x}, names[0]))
    print(json.dumps(res, ensure_ascii=False))
