from util import *

with open("data/非遗数据.csv", "r", encoding="utf-8") as f:
    data = f.readlines()
    with open("data/preprocessed.csv", "w", encoding="utf-8") as fout:

        header = data[0].split(",")
        for i in range(len(header)):
            header[i] = header[i].strip()[1:-1]
        # 为header添加一个第一列"ID"，和现在的最后一列前的三列"省份","民族","关键词"并写入文件
        header.insert(0, "UID")
        header.insert(-1, "省份")
        header.insert(-1, "民族")
        header.insert(-1, "关键词")
        newHeader = ",".join(header) + "\n"
        fout.write(newHeader)

        with open('data/desc.txt', 'w', encoding='utf-8') as fdesc:
            mp = {}
            for line in data[1:]:
                split_line = line.split(",")
                # 将第九列后面的全部重新合并，因为有些文本中有逗号，导致分割错误，合并后删除原来的列
                desc_str = "，".join(split_line[8:]).strip()[1:-1]
                if desc_str == '' or desc_str in mp:
                    continue
                fdesc.write(desc_str + '\n')
                mp[desc_str] = 1


        for line in data[1:]:
            split_line = line.split(",")
            # 将第九列后面的全部重新合并，因为有些文本中有逗号，导致分割错误，合并后删除原来的列
            split_line[8] = "，".join(split_line[8:])
            split_line = split_line[:9]
            for i in range(len(split_line)):
                split_line[i] = split_line[i].strip()[1:-1]
            # 添加一个第一列的ID，和"省份","民族","关键词"三列并写入文件
            split_line.insert(0, str(data.index(line)))
            split_line.insert(-1, getProvince(split_line[7]))
            split_line.insert(-1, getEthnic(split_line[10], split_line[3]))
            split_line.insert(-1, getKeyword(split_line[11], 20, 'data/banned_words.txt'))

            newLine = ",".join(split_line) + "\n"  # 在每行末尾添加换行符
            fout.write(newLine)  # 使用write函数逐行写入数据