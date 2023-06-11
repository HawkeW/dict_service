
import logging
import multiprocessing
import re
from bs4 import BeautifulSoup
from multiprocess import process_array_with_progress, process_data

from parsers.caption import WordCaption


class ColinsCnParser:
  def __init__(self, word, html):
    self.word = word
    self.pronUk = []
    self.pronUs = []
    self.captions = []
    self.link = ''
    self.parse(html)

  # 解析html
  def parse(self, html):
    match = re.search(r'@@@LINK=(.+)', html)
    if match is None:
      soup = BeautifulSoup(html, 'html.parser')
      if bool(soup):
        self.getProns(soup)
        self.captions = self.getCaptions(soup)
    else:
      link = match.group(1)
      self.link = link
        # print("No match found. html", html) 

  # 获取link指向单词
  def getLinkLink(self, doc):
    res = re.search(r'@@@LINK=(\w+)', doc)
    if bool(res): return res.group(1)
    return ''

  # 获取所有发音元素
  def getProns(self, doc):
    # 查找所有的<span class="pron">元素
    pron_span = doc.find('span', class_='pron')
    # 提取pronUk和pronUs的内容
    pronUk = []
    pronUs = []
    if pron_span is not None:
      # 获取所有的<span class="pron type_uk">元素
      pron_uk_spans = pron_span.find_all('span', class_='pron type_uk')

      # 获取所有的<span class="pron type_us">元素
      pron_us_spans = pron_span.find_all('span', class_='pron type_us')

      for span in pron_uk_spans:
          pron = span.text.strip()
          pron_a = span.find_parent('a')
          mp3 = ''
          if pron_a is not None:
            mp3 = pron_a['href']
          pronUk.append({'pron': pron, 'mp3': mp3})

      for span in pron_us_spans:
          pron = span.text.strip()
          pron_a = span.find_parent('a')
          mp3 = ''
          if pron_a is not None:
            mp3 = pron_a['href']
          pronUs.append({'pron': pron, 'mp3': mp3})
    
    self.pronUk = pronUk
    self.pronUs = pronUs

  # 获取词性及对应释义数组
  def getCaptions(self, doc):
    # 查找所有的<div class="example">元素
    example_div_list = doc.find_all('div', class_='example')
    captionList = []
    if example_div_list is not None:
      for example in example_div_list:
        caption = WordCaption('','','','',[])
        caption_div = example.find('div', class_="caption")
        if caption_div is None: continue
        caption_span_st = caption_div.find('span', class_="st")
        # 单词词性
        if caption_span_st is not None:
          caption.st = caption_span_st.text.strip()
          caption.stCn = caption_span_st["title"] if caption_span_st.has_attr("title") else ''
        # 单词释义
        caption_def_cn = caption_div.find('span', class_="def_cn cn_after")
        if caption_def_cn is not None:
          caption.defCn = caption_def_cn.text.strip()
          caption.defEn = caption_def_cn.previous_sibling.strip()

        # ul元素为例句元素
        sentence_list_ul = example.find('ul')
        if sentence_list_ul is not None:
          # 存在例句
          sentence_li_list = sentence_list_ul.find_all('li', recursive=False)
          for sentence_li in sentence_li_list:
            # 英文例句
            sentenceEn = sentence_li.contents[0].text.strip()
            # 英文例句翻译
            sentenceCn = sentence_li.contents[1].text.strip()
            caption.sentences.append({"en": sentenceEn,"cn": sentenceCn})

          captionList.append(caption)
    return captionList


def parseColins():
  wordList = []

  with open('D:\BaiduNetdiskDownload\Dictionary\Collins\CollinsCOBUILDOverhaul V 2-30.html', 'r', encoding='UTF-8') as f:
    dmx_doc = f.read()
  
    list = dmx_doc.split('</>')
    
    # processed_list = process_array_with_progress(list[1:20], parseColinsPartialListWithProcess, 2)
    wordList = parseColinsList(list)
 
  return wordList

# 解析数组数据
def parseColinsPartialListWithProcess(list, processIndex, processQueue):
  # 创建日志记录器
  logger = multiprocessing.get_logger()
  logger.setLevel(logging.INFO)
  try:
    # 创建文件处理器并设置日志格式
    file_handler = logging.FileHandler(f"process_{processIndex}.log")
    formatter = logging.Formatter("%(asctime)s - %(levelname)s - %(message)s")
    file_handler.setFormatter(formatter)

    # 将文件处理器添加到日志记录器
    logger.addHandler(file_handler)

    # 在子进程中记录日志
    logger.info(f"Processing data in process {processIndex}")

    # 在这里定义你的处理逻辑
    words = parseColinsList(list) 
  except Exception as e:
      # 记录异常信息到日志
      logging.exception("Error occurred during data processing")
  
  # 示例日志记录
  logger.info(f"processed {len(words)}, process_id: {processIndex}")
  
  # 向进程队列发送进度信息
  processQueue.put(processIndex)

def parseColinsList(list):
  wordList = []
  for item in list:
    res = re.search(r'\n(.+)\n(.+)', item)
    if bool(res):
      wordList.append(ColinsCnParser(res.group(1), res.group(2)))
      print(f'parsed {len(wordList)}')
  return wordList