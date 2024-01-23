Проект запускается в терминале командой __***"docker-compose up***"__

Сервер слушает __80__ порт локалхоста .

http://127.0.0.1:80/users/

__addUser__ - добавление нового пользователя в базу данных(БД) . Получает значения "*name*", "*surname*", "*patronymic*" в теле POST-запроса в формате *JSON* , по значению *"name"* обращается на сторонние ресурсы для получения "*age*", "*gender*", "*nationality*", после добавляет в базу данных новую запись со всеми полученныими параметрами, присваивая уникальный ID это записи . Возвращает данные добавленного пользователя в формате *JSON* .

__getUser__ - чтение данных одного пользователя по *ID* из БД . Получает значение *ID* из URL GET-запроса , по полученному значению *ID* считывает данные пользователя из БД . Возвращает данные пользователя в формате *JSON* .

__getUsers__ - чтение данных одного или нескольких пользователей по определенному фильтру из БД . Получает критерии поиска из queryString GET-запроса в формате "*?attr=value*" , где *attr* - любое поле для фильтра, *value* - значение поиска . Можно добавлять несколько критериев поиска , перечислив их в стоке запроса через символ "*&*" . Если отправить пустой запрос , то вернутся все имеющиеся записи из БД . Возвращает данные пользователей в формате *JSON* .

__deleteUser__ - удаление данных пользователя по *ID* из БД . Получает значение *ID* из URL DEL-запроса , по полученному значению ID удаляет данные пользователя из БД . Возвращает данные удаленного пользователя в формате *JSON* .

__updateUser__ - изменение данных пользователя по *ID* из БД . Получает значение *ID* из URL PUT-запроса , новые данные пользователя должны быть перечисленны в теле запроса в формате *JSON* . Возвращает данные измененного пользователя в формате *JSON* .