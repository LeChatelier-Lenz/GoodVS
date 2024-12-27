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
    page = SessionPage()
    # 访问某一页的网页
    page.get(f'https://search.suning.com/{key_string}/')
    print(page)
    try:
        category = page.ele('.class-relevant').ele('.r-name').text
    except Exception as e:
        category = '未知'
        print(e)
    print(category)
    links = page.eles('.product-box ')
    # 遍历所有<a>元素
    result = []
    for link in links:
        # 打印链接信息
        # print(link.html)
        img_url = link.ele('.res-img').ele('.img-block').child().child().attr('src')
        # print(img_url)
        info = link.ele('.res-info')
        price = info.ele('.def-price').text
        title = info.ele('.title-selling-point').child().text
        result.append({
            'imgUrl': img_url,
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
    words = jieba.lcut(text)
    return words


if __name__ == '__main__':
    input_str = sys.argv[1]
    input_arr = segment_text(input_str)
    r = get_suning(input_arr)
    # for i in r:
    #     print(i)
