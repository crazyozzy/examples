import sqlite3 as sq
# import time
import telebot
import speech_recognition as sr
# import inspect
import ffmpeg
# from pydub import AudioSegment
from telebot import types

token = '<your_token>'
bot = telebot.TeleBot(token)

HELP = '''
Список доступных команд:
* show  - напечать все задачи на заданную дату
* add - добавить задачу
* todo - показать невыполненные задачи
*
* help - Напечатать help
'''

crossIcon = u"\u274C"

def make_keyboard_complete(task_list):
    markup = types.InlineKeyboardMarkup()
    for id, complete, date, task in task_list:
        markup.add(types.InlineKeyboardButton(text=crossIcon, callback_data=f"complete_task {id}"),
                   types.InlineKeyboardButton(text=f"{date} {task}", callback_data=f"complete_task {id}"))
    return markup


def gen_markup():
    markup = types.InlineKeyboardMarkup()
    markup.row_width = 2
    markup.add(types.InlineKeyboardButton("todo", callback_data="cb_yes"),
               types.InlineKeyboardButton("No", callback_data="cb_no"))
    return markup


def add_to_db(chat_id, task, date):
    with sq.connect('easyplanner.db') as conn:
        cur = conn.cursor()
        cur.execute(f"INSERT INTO tasks (chat_id, task, date) VALUES('{chat_id}', '{task}', '{date}')")

def spoken_text_handler(input, chat_id):
    input = str(input.split())
    date: str = input[1:4]
    date[1].replace(date[1].lower())
    try:
        date[0].replace(int(date[0]))
    except Exception:
        bot.send_message(chat_id, "Нет такого числа")
        return 1
    if 1 <= date[0] <= 9:
        date[0].replace(f'0{date[0]}')
    elif 10 <= date[0] <= 31:
        date[0].replace(f'{date[0]}')
    else:
        print('Нет такого числа')
        return 1

    if date[1] == 'январь' or date[1] == 'января':
        date[1].replace('01')
    elif date[1] == 'февраль' or date[1] == 'февраля':
        date[1].replace('02')
    elif date[1] == 'март' or date[1] == 'марта':
        date[1].replace('03')
    elif date[1] == 'апреля' or date[1] == 'апреля':
        date[1].replace('04')
    elif date[1] == 'май' or date[1] == 'мая':
        date[1].replace('05')
    elif date[1] == 'июнь' or date[1] == 'июня':
        date[1].replace('06')
    elif date[1] == 'июля' or date[1] == 'июль':
        date[1].replace('07')
    elif date[1] == 'август' or date[1] == 'августа':
        date[1].replace('08')
    elif date[1] == 'сентябрь' or date[1] == 'сентября':
        date[1].replace('09')
    elif date[1] == 'октябрь' or date[1] == 'октября':
        date[1].replace('10')
    elif date[1] == 'ноябрь' or date[1] == 'ноября':
        date[1].replace('11')
    elif date[1] == 'декабрь' or date[1] == 'декабря':
        date[1].replace('12')
    else:
        print('Месяц не определен')
        return 1
    date.append(f'{date[0]}.{date[1]}.{date[2]}')

    if input[0].lower() == 'создать':
        add_to_db(chat_id, input[4:].join(), date)



@bot.message_handler(commands=['help'])
def help(message):
    bot.send_message(message.chat.id, HELP)


@bot.message_handler(commands=['show'])
def show_tasks(message):
    date = message.text.split()[1].lower()
    with sq.connect('easyplanner.db') as conn:
        cur = conn.cursor()
        cur.execute(f"SELECT task, complete FROM tasks WHERE date='{date}' AND chat_id='{message.chat.id}'")
    # if date in todos:
    #     tasks = ''
    #     for task in todos[date]:
    #         tasks += f'[ ] {task}\n'
    # else:
    #     tasks = 'Такой даты нет'
    tasks = ''
    for row in cur.fetchall():
        if row[1] == 0:
            tasks += f'[  ] {row[0]}\n'
        else:
            tasks += f'[X] {row[0]}\n'
    bot.send_message(message.chat.id, tasks)


@bot.message_handler(commands=['todo'])
def show_tasks(message):
    # date = message.text.split()[1].lower()
    with sq.connect('easyplanner.db') as conn:
        cur = conn.cursor()
        cur.execute(f"SELECT task, date FROM tasks WHERE chat_id='{message.chat.id}' AND complete=0")
    # if date in todos:
    #     tasks = ''
    #     for task in todos[date]:
    #         tasks += f'[ ] {task}\n'
    # else:
    #     tasks = 'Такой даты нет'
    tasks = ''
    for row in cur.fetchall():
        tasks += f'[  ] {row[1]} {row[0]}\n'
    bot.send_message(message.chat.id, tasks)


@bot.message_handler(commands=['add'])
def add(message):
    _, date, tail = message.text.split(maxsplit=2)
    task = ' '.join([tail])
    with sq.connect('easyplanner.db') as conn:
        cur = conn.cursor()
        cur.execute(f"INSERT INTO tasks (chat_id, task, date) VALUES('{message.chat.id}', '{task}', '{date}')")
    bot.send_message(message.chat.id, f'Задача {task} добавлена на дату {date}')


@bot.message_handler(commands=['complete'])
def complete(message):
    with sq.connect('easyplanner.db') as conn:
        cur = conn.cursor()
        cur.execute(f"SELECT id, complete, date, task FROM tasks WHERE chat_id='{message.chat.id}' AND complete=0")
    # print(cur.fetchall())
    bot.send_message(chat_id=message.chat.id,
                     text="Here are tasks",
                     reply_markup=make_keyboard_complete(cur.fetchall()),
                     parse_mode='HTML')


@bot.message_handler(commands=['showall'])
def show_all(message):
    with sq.connect('easyplanner.db') as conn:
        cur = conn.cursor()
        cur.execute(f"SELECT complete, date, task FROM tasks WHERE chat_id='{message.chat.id}'")
    tasks = ''
    for complete, date, task in cur.fetchall():
        if complete == 0:
            tasks += f'[  ] {date} {task}\n'
        else:
            tasks += f'[X] {date} {task}\n'
    bot.send_message(message.chat.id, tasks)


@bot.message_handler(commands=['test'])
def handle_command_adminwindow(message):
    bot.send_message(chat_id=message.chat.id,
                     text="Here are the values of stringList",
                     reply_markup=gen_markup(),
                     parse_mode='HTML')


@bot.message_handler(content_types=['voice'])
def handle_audio(message):
    r = sr.Recognizer()
    file_info = bot.get_file(message.voice.file_id)
    downloaded_file = bot.download_file(file_info.file_path)
    with open('input_voice.ogg', 'wb') as new_file:
        new_file.write(downloaded_file)
    ffmpeg.input('input_voice.ogg').audio.output('output_voice.flac').overwrite_output().run()
    with sr.AudioFile('output_voice.flac') as output:
        voice = r.record(output)
    # print(r.recognize_sphinx(message.voice))
    rec_text = r.recognize_google(voice, language='ru-RU')
    bot.send_message(message.chat.id, rec_text)
    spoken_text_handler(rec_text, message.chat.id)


@bot.callback_query_handler(func=lambda call: call.data == "cb_yes")
def callback_query(call):
    # bot.answer_callback_query(call.id, "Answer is Yes")
    bot.send_message(call.from_user.id, 'Your answer is YES!')
    # elif call.data == "cb_no":
    #     bot.answer_callback_query(call.id, "Answer is No")


@bot.callback_query_handler(func=lambda call: call.data.split()[0] == "complete_task")
def complete_task(call):
    with sq.connect('easyplanner.db') as conn:
        cur = conn.cursor()
        cur.execute(f"UPDATE tasks SET complete=1 WHERE id={call.data.split()[1]}")
    bot.answer_callback_query(call.id, f"Task deleted")


with sq.connect('easyplanner.db') as conn:
    cur = conn.cursor()
    cur.execute("""CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id TEXT,
    task TEXT,
    date TEXT,
    complete INT DEFAULT 0 NOT NULL
    )""")

# audio = ffmpeg.input('input_voice.ogg').audio
# stream = ffmpeg.output(audio, 'output_voice.flac')
# ffmpeg.run(stream, capture_stdout=True, capture_stderr=True)

bot.polling(none_stop=True)
