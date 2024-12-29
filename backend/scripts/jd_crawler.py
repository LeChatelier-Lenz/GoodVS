import json
import sys
import jieba

from DrissionPage import Chromium


# def pure_jd_login():
#     page = Chromium().latest_tab
#     page.get("https://www.jd.com/")
#     # 输出原网站的地址
#     # print(page.title)
#     if page.ele('.link-login') and not page.ele('.nikename'):
#         jd_login(page)
#     return page


def jd_login():
    page = Chromium().new_tab()
    page.get("https://www.jd.com/")
    page.ele('.link-login').click()
    # 输出当前网站的标题
    qr_img_url = page.ele('.qrcode-img').child().attr('src')
    print(qr_img_url)
    print('请扫码登录')
    # 等待扫码，一直到.qrcode-img元素消失
    while page.ele('.qrcode-img'):
        pass
    print('扫码成功')
    cookies = page.cookies(all_domains=True).as_dict()
    f_cookies = open('./tmp/jd_cookies.json', 'w', encoding='utf-8')
    f_cookies.write(json.dumps(cookies, ensure_ascii=False))
    f_cookies.close()
    return cookies


def get_JD(keywords):
    try:
        f_cookies = open('./tmp/jd_cookies.json', 'r', encoding='utf-8')
        cookies = json.load(f_cookies)
        f_cookies.close()
    except Exception as e:
        print(e)
        print('未找到cookies文件')
        cookies = jd_login()
    key_string = " ".join(keywords)
    page = Chromium().new_tab()
    # page.get("https://www.jd.com/")
    # 输出原网站的地址
    # print(page.title)
    # if page.ele('.link-login') and not page.ele('.nikename'):
    #     jd_login(page)
    try:
        page.get(f"https://search.jd.com/Search?keyword={key_string}&enc=utf-8")
        ele = page.ele(".J_selectorLine s-category").ele(".sl-key")
    except Exception as e:
        # cookies失效，重新登录
        print('登录状态失效')
        print(e)
        try:
            jd_login()
            page.get(f"https://search.jd.com/Search?keyword={key_string}&enc=utf-8")
            ele = page.ele(".J_selectorLine s-category").ele(".sl-key")
        except Exception as reLoginError:
            print(reLoginError)
            print('登录失败')
            exit(0)
    if ele:
        category = ele.text
    else:
        category = '未知种类'
    links = page.eles('.gl-i-wrap')
    if not links:
        print('未找到商品')
    result = []
    for link in links[:5]:
        # print(link.html)
        img_url = link.ele('.p-img').child().child().attr('src')
        # count = 0
        while img_url == '' or img_url is None:
            page.actions.scroll(5, 0)
            img_url = link.ele('.p-img').child().child().attr('src')
        price = link.ele('.p-price').child().text
        # if price.find('已补贴'):
        #     price = price.split('已补贴')[0]
        # elif price.find("到手价"):
        #     price = price.split('到手价')[0]
        if price.find('¥'):
            print(price)
            price = price.split('￥')[1]
        title = link.ele('.p-name p-name-type-2').child().text
        result.append({
            'name': key_string,
            'img_url': img_url,
            'price': float(price),
            'title': title,
            'platform': '京东',
            'category': category
        })
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


def conclude_category(key_string):
    f = open('./tmp/category.txt', 'r', encoding='utf-8')
    str_result = f.read()
    f.close()
    category = str_result.split('/')
    # 根据key_string的内容，判断其属于哪个类别,返回类别
    # 并不一定有匹配的上的内容，所以要调用库


if __name__ == '__main__':
    # 脚本使用说明
    # python jd_crawler.py <options> <keywords>
    # options: r(后端重复查询，已经过分词) f(第一次查询，未经过分词) login(登录)
    # keywords: 搜索关键字
    choice = sys.argv[1]
    if choice == 'r':
        input_arr = sys.argv[2:]
    elif choice == 'f':
        input_arr = segment_text(sys.argv[2:])
        print("分词结果:", input_arr)
    elif choice == 'login':
        # 登录模式
        jd_login()
        exit(0)
    else:
        print('参数错误')
        exit(0)
    if len(input_arr) == 0:
        print('未输入搜索关键字')
        exit(0)
    r = []
    # 非登录模式
    r_JD = get_JD(input_arr)
    r.extend(r_JD)
    f = open('./tmp/jd_result.json', 'w', encoding='utf-8')
    f.write(json.dumps(r, ensure_ascii=False))
    f.close()

