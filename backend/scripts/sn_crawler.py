import json
import sys
import jieba

from DrissionPage import Chromium
from DrissionPage import SessionPage


def get_taobao(keywords):
    pass


def get_SN(keywords):
    # 创建页面对象
    key_string = " ".join(keywords)
    page = Chromium().new_tab()
    # 访问某一页的网页
    page.get(f'https://search.suning.com/{key_string}/')
    # print(page)
    try:
        category = page.ele('.class-relevant').ele('.r-name').text
    except Exception as out_e:
        try:
            category = page.ele('.result-right').text
        except Exception as e:
            category = '未知种类'
            print(e)
        print(out_e)
    print(category)
    links = page.eles('.product-box ')
    # 遍历所有<a>元素
    result = []
    for link in links[:5]:
        img_url = link.ele('.res-img').ele('.img-block').child().child().attr('src')
        price = link.ele('.def-price').text
        while price == '':
            page.actions.scroll(5, 0)
            price = link.ele('.def-price').text
        if price.find('¥') == -1:
            price = float(price.split('到')[0])
        else:
            price = float(price.split('¥')[1].split('到')[0])
        title = link.ele('.title-selling-point').child().text
        result.append({
            'name': key_string,
            'img_url': img_url,
            'price': price,
            'title': title,
            'platform': '苏宁',
            'category': category
        })
        # print(img_url)
    page.close()
    return result


# SegmentText 函数用于对输入的文本进行分词，并去除停用词和排序
def segment_text(text_segs):
    # 使用 jieba 进行分词，`cut` 方法默认为精确模式，`HMM=True` 启用 HMM 模型
    words = jieba.cut(" ".join(text_segs), HMM=True)
    words = list(words)  # 将迭代器转换为列表
    # print("搜索模式:", words)

    if not words:
        raise ValueError("jieba cut failed")

    # 去除停用词
    words = remove_stop_words(words)

    # 按照一定规则对分词结果进行排序
    words.sort()
    return words


# RemoveStopWords 函数用于去除停用词
def remove_stop_words(words):
    stop_words = {"的", "了", "和", "是", "就", "都", "而", "及", "与", "或", "于", "之", "从", "也", "在"}
    res = []

    for w in words:
        # 判断是否包含停用词，如果包含则去除停用词
        for s in stop_words:
            if s in w:
                if w != s:  # 保留非完全匹配的词
                    res.append(w.replace(s, ""))  # 去除停用词
                break
        else:
            # 如果不包含停用词，直接保留该词
            res.append(w)
    return res


if __name__ == '__main__':
    choice = sys.argv[1]
    r = []
    if choice == 'r':
        input_arr = sys.argv[2:]
    elif choice == 'f':
        input_arr = segment_text(sys.argv[2:])
        print("分词结果:", input_arr)
    else:
        print('参数错误')
        exit(0)
    if len(input_arr) == 0:
        print('未输入搜索关键字')
        exit(0)
    r_SN = get_SN(input_arr)
    r.extend(r_SN)
    # 方法2：将结果写入文件(json)
    # import json
    f = open('./tmp/sn_result.json', 'w', encoding='utf-8')
    f.write(json.dumps(r, ensure_ascii=False))
    f.close()

    # 方法3：将结果打印到控制台
    # print(r)
