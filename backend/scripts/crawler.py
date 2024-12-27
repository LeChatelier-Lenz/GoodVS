import json
import sys
import jieba

from DrissionPage import Chromium
from DrissionPage import SessionPage


def get_jingdong(kerwords):
    pass


def get_taobao(keywords):
    pass


def get_suning(keywords):
    # 创建页面对象
    key_string = " ".join(keywords)
    page = Chromium().latest_tab
    # 访问某一页的网页
    page.get(f'https://search.suning.com/{key_string}/')
    # print(page)
    try:
        category = page.ele('.class-relevant').ele('.r-name').text
    except Exception as e:
        try:
            category = page.ele('.result-right').text
        except Exception as e:
            category = '未知种类'
            print(e)
        print(e)
    print(category)
    links = page.eles('.product-box ')
    # 遍历所有<a>元素
    result = []
    for link in links[:3]:
        # 打印链接信息
        # print(link.html)
        img_url = link.ele('.res-img').ele('.img-block').child().child().attr('src')
        # special = link.ele('.price-arrival')
        # print(special)
        # print(img_url)
        # info = link.ele('.res-info')
        # print(info.html)
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
            'name': " ".join(keywords),
            'img_url': img_url,
            'price': price,
            'title': title,
            'platform': '苏宁',
            'category': category
        })
        # print(img_url)
    return result


# # 连接浏览器并获取一个MixTab对象
# tab = Chromium().latest_tab
# # 访问网址
# tab.get('https://gitee.com/explore/all')
# # 切换到收发数据包模式
# tab.change_mode()
# # 获取所有行元素
# items = tab.ele('.ui relaxed divided items explore-repo__list').eles('.item')
# # 遍历获取到的元素
# for item in items:
#     # 打印元素文本
#     print(item('t:h3').text)
#     print(item('.project-desc mb-1').text)
#     print()


def segment_text(text):
    # 使用 jieba 对用户输入进行分词
    if type(text) is str:
        words = jieba.lcut(text)
    elif type(text) is list:
        words = []
        for i in text:
            words.extend(jieba.lcut(i))
    else:
        words = []
    return words


if __name__ == '__main__':
    input_str = sys.argv[1:]
    input_arr = segment_text(input_str)
    if len(input_arr) == 0:
        print('未输入搜索关键字')
    # print(input_arr)
    r = get_suning(input_arr)
    # 方法1：将结果写入文件(txt)
    # f = open('result.txt', 'w', encoding='utf-8')
    # f.write('[')
    # for i in r:
    #     f.write(str(i) + ',')
    # f.write(']')
    # f.close()
    # 方法2：将结果写入文件(json)
    # import json
    f = open('result.json', 'w', encoding='utf-8')
    f.write(json.dumps(r, ensure_ascii=False))
    f.close()

    # 方法3：将结果打印到控制台
    # print(r)
