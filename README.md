# Система тестирования 
WEB сервис "Система тестирования" для создания, прохождения и просмотра результатов тестов. Система с пользовательским API
для взаимодействия личных программных продуктов клиента с серверами системы тестирования. 
____
# Технологии
- Backend
  - Go 
  - PostgreSQL
- Frontend
  - Scss
  - TypeScript
  - NodeJS
  - React
 ____
# Подготовка к запуску локального сервера 
Чтобы склонировать проект себе на пк, вы должны зайти через консоль в папку с проектами go и выполнить команду для клонирования.
``` 
git clone https://github.com/Zefirnutiy/sweet_potato.git 
```

После этого перейдите на рабочую ветку develop
``` 
git checkout develop
```

Далее создайте базу данных с произвольным именем и выполните в ней SQL код из файла **main.sql**, который вы найдёте по пути **.\backend\db** . <br>
После создания базы данных, в корне проекта создайте файл **settings.cfg** и заполните его данными о сервере и базе данных.
```
{
    "Host": "127.0.0.1", # Локальный IP
    "Port": ":8080",     # Порт на котором будет запускаться сервер
    "Database":{
        "DbHost": "",    # IP на котором запущена база данных
        "DbPort": "",    # Порт на котором запущена база данных
        "DbUser": "",    # Имя владельца базы данных
        "DbPass": "",    # Пароль от базы данных
        "DbName": ""     # Имя базы данных
    }
}
```
Зайдите через консоль в папку с react по пути **.\frontend\client-ts** и выполните команду для докачки node_modules.
```
npm i
```
____
# Запуск локального сервера
Наш проект работает на 2-х серверах. На go и react, так что придётся запускать их оба для полноценной работы сайта. <br>
**Сервер go -** Запустите файл **main.go**, который в папке **./backend**
``` 
go run main.go 
```
**Сервер react -** Зайдите через консоль в папку с react по пути **.\frontend\client-ts** и выполните команду.
```
npm run start
```
После запуска go и react, сервер полностью работает.