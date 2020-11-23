## Домашнее задание №10 «Оптимизация программы»

Вам дан исходный код функции `GetDomainStat(r io.Reader, domain string)`, которая:
* читает построчно из `r` пользовательские данные вида
```text
{"Id":1,"Name":"Howard Mendoza","Username":"0Oliver","Email":"aliquid_qui_ea@Browsedrive.gov","Phone":"6-866-899-36-79","Password":"InAQJvsq","Address":"Blackbird Place 25"}
{"Id":2,"Name":"Brian Olson","Username":"non_quia_id","Email":"FrancesEllis@Quinu.edu","Phone":"237-75-34","Password":"cmEPhX8","Address":"Butterfield Junction 74"}
{"Id":3,"Name":"Justin Oliver Jr. Sr.","Username":"oPerez","Email":"MelissaGutierrez@Twinte.gov","Phone":"106-05-18","Password":"f00GKr9i","Address":"Oak Valley Lane 19"}
```
(осторожно, в отличие от конкретной строки файл целиком не является валидным JSON);
* подсчитывает количество email-доменов пользователей на основе домена первого уровня `domain`.

Например, для данных, представленных выше:
```text
GetDomainStat(r, "com") // {}
GetDomainStat(r, "gov") // {"browsedrive": 1, "twinte": 1}
GetDomainStat(r, "edu") // {"quinu": 1}
```

Для большего понимания см. исходный код и тесты.

**Необходимо оптимизировать программу таким образом, чтобы она проходила все тесты.**

Нельзя:
- изменять сигнатуру функции `GetDomainStat`;
- удалять или изменять существующие юнит-тесты.

Можно:
- писать любой новый необходимый код;
- удалять имеющийся лишний код (кроме функции `GetDomainStat`);
- добавлять юнит-тесты.

**Обратите внимание на запуск TestGetDomainStat_Time_And_Memory**
```bash
go test -v -count=1 -timeout=30s -tags bench .
```

Здесь используется билд-тэг bench, чтобы отделить обычные тесты от тестов производительности.

### Критерии оценки
- Пайплайн зелёный и нет попытки «обмануть» систему - 4 балла
- Добавлены юнит-тесты - до 3 баллов
- Понятность и чистота кода - до 3 баллов

### Частые ошибки
- Работа с сырыми байтами, нахождение позиции `"Email"` и пр. вместо ускорения анмаршалинга более поддерживаемыми и понятными средствами.

#### Зачёт от 7 баллов
