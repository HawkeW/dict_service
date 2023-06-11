import csv

from database.words import closeDb, insert_word, update_link
from parsers.colinsCn import parseColins 

if __name__ == '__main__':
  # 解析词典
  wordList = parseColins()

  # 定义CSV文件路径
  csv_file = 'words.csv'

  # 写入CSV文件
  # with open(csv_file, 'w', newline='', encoding='utf-8') as f:
  #     # 创建CSV写入器
  #     writer = csv.writer(f)

  #     # 写入表头
  #     writer.writerow(['ID', '单词', '注解', '英式发音', '美式发音', '关联单词'])

  #     # 写入数据行
  #     for index, w in enumerate(wordList):
  #       wordCaption = [caption.__dict__ for caption in w.captions]
  #       writer.writerow([ index, w.word, wordCaption, w.pronUk, w.pronUs, w.link])
  # print("CSV文件写入完成。")

  for index, word in enumerate(wordList):
    insert_word(word)
    print(f'insertWord {index}/{len(wordList)}')

  for index, word in enumerate(wordList):
    update_link(word)
    print(f'updateLink {index}/{len(wordList)}')

  closeDb()