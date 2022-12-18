# Основы Ansible

**Ansible** — комплекс ППО для управления инфраструктурной составляющей ваших систем и развёртыванием приложений:
* прост в использовании
* не требует установки агентов
* идемпотентен
* легко расширяем

Внутри Ansible существуют следующие понятия:
* Playbook
* Play
* Role
* Task
* Handlers
* Inventory
* Group vars
* Facts
* Templates
* Collections

## Основные команды Ansible
* ansible – определение и запуск playbook из одного task на наборе hosts
* ansible-playbook — запуск полноценного playbook
* ansible-vault — шифрование хранилища методом AES256
* ansible-galaxy — скачивание roles и collections
* ansible-lint — используется для проверки синтаксиса
* ansible-console — консоль для запуска tasks
* ansible-config — просмотр и управление конфигурацией ansible
* ansible-doc — просмотр документации plugins
* ansible-inventory — просмотр информации о hosts из inventory
* ansible-pull — скачивание playbook и запуск на localhost
* ansible-test — тестирование collections

## Базовые команды
`ansible -m <имя модуля, например ping> <имя удаленного хоста>` - запуск adhoc команды на удаленном хосте, например `ansible -m ping localhost`
`ansible -m <имя модуля> -i <путь до inventory файла> <имя группы хостов для запуска>` - запуск adhoc команды на группе удаленных хостов из inventory файла, например `ansible -m ping -i inventory.yml all`
`ansible-playbook -i <путь до inventory файла> <путь до файла с playbook>` - запуск playbook на хостах из inventory файла, например `ansible-playbook -i inventory.yml site.yml`
`ansible-inventory -i <путь до inventory файла> --graph <имя группы>` - отображение графа inventory файла по группе хостов, например `ansible-inventory -i inventory.yml --graph all`
`ansible-inventory -i <путь до inventory файла> --list` - отображение всех переменных для всех хостов из inventory
`ansible-inventory -i <путь до inventory файла> --list <имя хоста>` - отображение всех переменных по определенному хосту
`ansible-doc <имя плагина>` - получение документации по указанному плагину
`ansible-vault create <путь до файла>` - создание нового vault файла
`ansible-vault edit <путь до файла>` - редактирование vault файла стандартным текстовым редактором
`ansible-vault view <путь до файла>` - отображение содержимого vault файла
`ansible-vault rekey <путь до файла>` - смена пароля vault файла
`ansible-vault encrypt <путь до файла>` - зашифровать файл с переменными
`ansible-vault decrypt <путь до файла>` - расшифровать файл

## Playbook
Ansible Playbook — набор plays, содержащих в себе roles и\или tasks, которые выполняются на указанных в inventory хостах с определёнными параметрами для каждого из них или для их групп.

Ansible Playbook — набор plays, содержащих в себе roles и\или tasks, которые выполняются на указанных в inventory хостах с определёнными параметрами для каждого из них или для их групп Playbook описывается на языке YAML.

Пример содержимого одного play в Playbook:
```yaml
---
- name: Try run Vector                            # Произвольное название play
  hosts: all                                      # Перечисление хостов
  tasks:                                          # Объявление списка tasks
    - name: Get Vector version                    # Произвольное имя для task
      ansible.builtin.command: vector --version   # Что и как необходимо сделать
      register: is_installed                      # Запись результата в переменную is_installed
    - name: Get RPM                               # Произвольное имя для второй task
      ansible.builtin.get_url:                    # Объявление использования модуля get_url, ниже указание его параметров
        url: "https://package.timber.io/vector/{{ vector_version }}/vector.rpm"
        dest: {{ ansible_user_dir }}/vector.rpm
        mode: 0755
      when:                                       # Условия при которых task будет выполняться
        - is_installed is failed
        - ansible_distribution == "CentOS"
```

## Role 
Role — группа tasks, нацеленная на выполнение действий, приводящих к единому результату
* Role выполняет список действий
* Список может состоять из одного действия
* Role может быть написана самостоятельно или скопирована из `galaxy` при помощи команды `ansible-galaxy`
* Какие role скачивать, лучше указать в `requirements.yml`
* Role хранят по умолчанию в директории `roles`, **у каждой role своя директория внутри**

Пример использования role в рамках play:
```yaml
---
- name: Try run Vector  # Произвольное название play
  hosts: all            # Перечисление хостов
  roles:                # Объявление списка roles
    - vector            # Вызов роли vector из директории с roles
```

## Inventory
Inventory — директория с файлом или группой файлов, в которых описано, на каких хостах нужно выполнять действия
* Inventory может быть описан в виде стандартного файла host.ini или при помощи yaml-структуры
* Лучшая практика — использование yaml inventory

Пример inventory-файла:
```yaml
---
prod:
  children:
    nginx:
      hosts:
        prod-ff-74669-02:
          ansible_host: 255.245.12.32
          ansible_user: prod
  children:
    application:
      hosts:
        174.96.45.23:

test:
  children:
    nginx:
      hosts:
        localhost:
          ansible_connection: local
```

## Group vars
Group vars — в общем понимании файлы с переменными для групп хостов или для всех хостов, указанных в **inventory**
* По умолчанию хранятся в директории `group_vars`
* Определение переменных для всех хостов происходит в директории `all`
* Определение переменных для групп из `inventory` **происходит в соответствующих им директориях**
* Файлы с переменными могут называться, основываясь на внутренней логике playbook, сами имена важны для пользователей

## Facts
Facts — сбор информации об удалённом хосте, включая сетевую информацию, информацию о системе, информацию о пользователе и прочее.
* Можно собирать данные об одном хосте и использовать эти данные для настройки другого хоста
* Факты собираются автоматически в начале проигрывания **play**
* **ansible <hostname> -m setup** — получить facts с hostname
* Facts хранятся в переменной **ansible_facts**
* Сбор facts можно принудительно выключить, вписав в play `gather_facts: false`

## Приоритеты переменных
Переменные могут определяться и переопределяться на многих уровнях в Ansible.

Уровни приоритизации от меньшего к большему:
* значения из командной строки (`-u username`)
* значения по умолчанию из **roles**
* значения из файла **inventory**
* значения из файлов **group_vars/all**
* значения из файлов **group_vars/{groupname}**
* переменные из **play**
* значения переменных **role** из **vars**
* экстрааргументы из командной строки (`-e “user=myuser”`)

Полный перечень приоритетов можно увидеть в официальной [документации](https://docs.ansible.com/ansible/latest/playbook_guide/playbooks_variables.html#id16)

## Vault
**Ansible Vault** — инструмент, позволяющий зашифровать переменные (AES256), скрыв чувствительные данные от общего использования.

Пример команд:
`ansible-vault create <путь до файла>` - создание нового vault файла
`ansible-vault edit <путь до файла>` - редактирование vault файла стандартным текстовым редактором
`ansible-vault view <путь до файла>` - отображение содержимого vault файла
`ansible-vault rekey <путь до файла>` - смена пароля vault файла
`ansible-vault encrypt <путь до файла>` - зашифровать файл с переменными
`ansible-vault decrypt <путь до файла>` - расшифровать файл

## Templates
Templates — инструмент, позволяющий создать кастомизированный конфигурационный файл на основе шаблона.

Для шаблонизации используют **Jinja**.
* Любой конфигурационный файл (даже без переменных внутри) может быть использован
* Шаблон должен иметь расширение j2

Шаблонизация напоминает форматирование строк:
```
"Привет, {name}!".format(name="Мир")

>>> Привет, Мир!
```

## Collections
Collections — способ распространения контента Ansible.

Включает в себя набор roles, modules, playbooks.
* Наименование состоит из **namespace.collections**
* Под **namespace** понимается, например, название компании или нечто объединяющее все **collections** для **namespace**
* Под **collections** понимается само название коллекции
* Создаются и публикуются при помощи **ansible-galaxy**

