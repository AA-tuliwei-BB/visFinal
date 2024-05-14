def getProvince(str):
    if "北京" in str:
        return "北京"
    elif "天津" in str:
        return "天津"
    elif "河北" in str:
        return "河北"
    elif "山西" in str:
        return "山西"
    elif "内蒙古" in str:
        return "内蒙古"
    elif "辽宁" in str:
        return "辽宁"
    elif "吉林" in str:
        return "吉林"
    elif "黑龙江" in str:
        return "黑龙江"
    elif "上海" in str:
        return "上海"
    elif "江苏" in str:
        return "江苏"
    elif "浙江" in str:
        return "浙江"
    elif "安徽" in str:
        return "安徽"
    elif "福建" in str:
        return "福建"
    elif "江西" in str:
        return "江西"
    elif "山东" in str:
        return "山东"
    elif "河南" in str:
        return "河南"
    elif "湖北" in str:
        return "湖北"
    elif "湖南" in str:
        return "湖南"
    elif "广东" in str:
        return "广东"
    elif "广西" in str:
        return "广西"
    elif "海南" in str:
        return "海南"
    elif "重庆" in str:
        return "重庆"
    elif "四川" in str:
        return "四川"
    elif "贵州" in str:
        return "贵州"
    elif "云南" in str:
        return "云南"
    elif "西藏" in str:
        return "西藏"
    elif "陕西" in str:
        return "陕西"
    elif "甘肃" in str:
        return "甘肃"
    elif "青海" in str:
        return "青海"
    elif "宁夏" in str:
        return "宁夏"
    elif "新疆" in str:
        return "新疆"
    elif "台湾" in str:
        return "台湾"
    elif "香港" in str:
        return "香港"
    elif "澳门" in str:
        return "澳门"
    else:
        return "-"

def getEthnic(desc, name):
    all_ethnic = ["汉族", "蒙古族", "回族", "藏族", "维吾尔族", "苗族", "彝族", "壮族", "布依族", "朝鲜族", "满族", "侗族", "瑶族", "白族", "土家族", "哈尼族", "哈萨克族", "傣族", "黎族", "傈僳族", "佤族", "畲族", "高山族", "拉祜族", "水族", "东乡族", "纳西族", "景颇族", "柯尔克孜族", "土族", "达斡尔族", "仫佬族", "羌族", "布朗族", "撒拉族", "毛南族", "仡佬族", "锡伯族", "阿昌族", "普米族", "塔吉克族", "怒族", "乌孜别克族", "俄罗斯族", "鄂温克族", "德昂族", "保安族", "裕固族", "京族", "塔塔尔族", "独龙族", "鄂伦春族", "赫哲族", "门巴族", "珞巴族", "基诺族"]
    # 统计每个民族出现的次数，取出现次数最多的民族，作为该文本的民族
    count = {}
    # 统计次数
    for i in all_ethnic:
        if i in name:
            return i
        count[i] = desc.count(i)
    # 返回次数最多的民族
    result = max(count, key=count.get)
    return result if count[result] >= 2 else "-"

import jieba.analyse
import jieba.posseg as pseg
from collections import Counter

printed = {}

def getKeywordByFrequency(description, num_keywords=5, banned_words_file=None):
    # 分词并进行词性标注
    words_with_flag = pseg.cut(description)
    # 使用 Counter 统计词频
    counter = Counter()
    # 选取词频最高的前 num_keywords 个词，要求词性为名词、动名词或动词
    for word, flag in words_with_flag:
        if len(word) != 1 and (flag.startswith('n') or flag.startswith('v') or flag == 'i' or flag == 't' or flag == 'a'):
            counter[word] += 1
    # 提取词频最高的关键词
    keywords = [word for word, _ in counter.most_common(num_keywords)]

    # 如果提供了被禁止词文件路径，则读取文件中的被禁止词
    if banned_words_file:
        with open(banned_words_file, 'r', encoding='utf-8') as f:
            banned_words = f.read().split()
        # 过滤掉被禁止词
        keywords = [word for word in keywords if word not in banned_words]
    return keywords

def getKeyword(description, num_keywords=5, banned_words_file=None):
    """
    使用词频和Textrank算法从中文描述中提取关键词，并过滤掉被禁止的词语

    参数：
    - description：中文描述的字符串
    - num_keywords：要提取的关键词数量，默认为 5
    - banned_words_file：包含被禁止词语的文件路径，默认为 None

    返回：
    一个包含关键词的字符串，关键词之间用空格分隔
    """
    # 使用词频提取前 num_keywords/4 个关键词
    keywords = getKeywordByFrequency(description, num_keywords//4, banned_words_file)

    # 使用extract_tags提取前 num_keywords/4 个关键词
    words_by_tfidf = jieba.analyse.extract_tags(description, topK=num_keywords//2, allowPOS=('n', 'vn', 'v', 't', 'a', 'i'))
    # 将没在keywords中的词加入keywords，直到keywords的长度达到num_keywords//2
    for word in words_by_tfidf:
        if word not in keywords:
            keywords.append(word)
        if len(keywords) >= num_keywords//2:
            break
    
    # 使用Textrank算法提取 剩余的关键词
    keywords_by_textrank = jieba.analyse.textrank(description, topK=num_keywords, allowPOS=('n', 'vn', 'v', 't', 'a', 'i'))
    # 将没在keywords中的词加入keywords，直到keywords的长度达到num_keywords
    for word in keywords_by_textrank:
        if word not in keywords:
            keywords.append(word)
        if len(keywords) >= num_keywords:
            break
    
    # 如果提供了被禁止词文件路径，则读取文件中的被禁止词
    if banned_words_file:
        with open(banned_words_file, 'r', encoding='utf-8') as f:
            banned_words = f.read().split()
        # 过滤掉被禁止词
        keywords = [word for word in keywords if word not in banned_words]
    
    # 如果关键词数量超过了 num_keywords，只保留前 num_keywords 个关键词
    keywords = keywords[:num_keywords]

    # 将关键词列表合并为一个字符串，关键词之间用空格分隔
    keywords_str = " ".join(keywords)
    return keywords_str