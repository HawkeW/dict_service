
import json
import sqlite3


# 数据库处理
# 创建连接和游标
conn = sqlite3.connect('word_database.db')
cursor = conn.cursor()

# 创建单词表
cursor.execute('''
    CREATE TABLE IF NOT EXISTS words 
        (id INTEGER PRIMARY KEY AUTOINCREMENT,
        word TEXT,
        colins_cn_id INTEGER,
        colins_en_id INTEGER);
''')

cursor.execute('''
    CREATE TABLE IF NOT EXISTS colins_cn 
        (id INTEGER PRIMARY KEY AUTOINCREMENT,
        word TEXT,
        link_word_id INTEGER,
        pron_uk TEXT,
        pron_us TEXT,
        captions TEXT);
''')

# 插入单词数据， 初始link_word为空
def insert_word(word):
    # 检查colins_cn是否存在相同单词
    cursor.execute('SELECT id FROM colins_cn WHERE word = ?', (word.word,))
    existing_colins_word = cursor.fetchone()
    colins_id = -1
    if existing_colins_word is None:
       caps = json.dumps([wd.__dict__ for wd in word.captions])
       pronUk = json.dumps(word.pronUk)
       pronUs = json.dumps(word.pronUs)
       cursor.execute('''
            INSERT INTO colins_cn (word, link_word_id, pron_uk, pron_us, captions)
            VALUES (?, ?, ?, ?, ?)
        ''', (word.word, -1, pronUk, pronUs, caps))
       colins_id = cursor.lastrowid
    else:
       colins_id, = existing_colins_word

    # 检查words是否存在相同单词
    cursor.execute('SELECT id, colins_cn_id FROM words WHERE word = ?', (word.word,))
    existing_word = cursor.fetchone()
    

    # 插入words表
    if existing_word is None:
        # 不存在相同单词，执行插入操作
        cursor.execute('''
            INSERT INTO words (word, colins_cn_id)
            VALUES (?, ?)
        ''', (word.word, colins_id))
        conn.commit()
    elif existing_word[1] is None:
        # 更新 colins_cn _id
        cursor.execute('''
            UPDATE words SET colins_cn_id = ? WHERE word = ?
        ''', (colins_id, word.word))
        conn.commit()

# 单词数据插入完成后，再更新每个word的link
def update_link(word):
    if bool(word.link) and bool(word.word):
      cursor.execute('SELECT id FROM colins_cn WHERE word = ?', (word.link,))
      result = cursor.fetchone()
      if result is not None:
        id, = result
        # 更新 link_word
        cursor.execute('UPDATE colins_cn SET link_word_id = ? WHERE word = ?', (id, word.word))
        conn.commit()

def closeDb():
   conn.close()