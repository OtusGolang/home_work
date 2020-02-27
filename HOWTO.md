## Процесс сдачи домашнего задания
Для сдачи ДЗ необходимо выполнить следующие действия (по порядку):
* Нажмите на зелёную кнопку **«Use this template»** на странице данного репозитория.
* Подождите пока для сгенерируется новый репозиторий на основе текущего.
* Подключите к новому репозиторию Travis CI:
1) https://travis-ci.com/
2) Sign in with GitHub
3) Activate all repositories using GitHub Apps
4) Only selected repositories -> Выбираете ваш новый
5) Approve & Install
* Склонируйте себе ваш репозиторий (например, я назвал свой `hw-test`):
```bash
$ git clone https://github.com/Antonboom/hw-test.git
```
* Создайте ветку **с именем таким же, как директория, где лежит ДЗ**. Это важно!
```bash
$ cd hw-test
$ git checkout -b hw01_hello_now
```
* Реализуйте код домашнего задания.
* Проверьте, что следующие команды завершаются успешно:
```bash
$ golangci-lint run ./...
$ go test -count=1 -race -gcflags=-l ./...
$ ./test.sh # При наличии
```
Это те команды, которые запускаются в CI (см. [.travis.yml](./.travis.yml)).
Дополнительно CI проверяет работоспособность `go get` для текущего модуля
(в нашем случае `hw01_hello_now`).
* Зафиксируйте изменения и запушьте ветку в репозиторий:
```bash
$ git commit -am "HW1 is completed"
$ git push origin hw01_hello_now
...
remote: Create a pull request for 'hw01_hello_now' on GitHub by visiting:
remote:      https://github.com/Antonboom/hw-test/pull/new/hw01_hello_now
```
* Как видно выше, GitHub предложит вам URL для создания пулл реквеста, пройдите по нему.

![pr_without_template](./img/pr_without_template.png)

* Допишите в конец URL параметр вида `&template=<имя_ветки>.md` и нажмите Enter -
PR обновится в соответствии с одним из [шаблонов](./.github/PULL_REQUEST_TEMPLATE).

![pr_with_template](./img/pr_with_template.png)

* Нажмите кнопку «Create pull request».

* Зайдите на страницу настроек веток репозитория (Settings -> Branches):
    * выбрать Default branch - master;
    * добавить новое правило (Branch protection rules -> Add rule):
        * Branch name pattern - `master`;
        * выставить галочку "Require status checks to pass before merging";
        * выставить галочку "Require branches to be up to date before merging";
        * выставить галочку "Travis CI - Pull Request";
        * выставить галочку "Include administrators";
        * нажать кнопку «Create».

* Скинуть ссылку на PR в чат с преподавателем в личном кабинете OTUS.
* Пройти ревью и **после одобрения пулл реквеста** вмержить PR в master
(у вас будет доступна кнопка «Merge request» и до момента ответа от преподавателя,
но стоит сначала дождаться апрува от него).
* Complete!

Если вы не хотите, чтобы CI запускался на каждый push в ветку, а работал
только при пулл реквесте, то снимите галочку "Build pushed branches"
на странице настроек репозитория в Travis CI
(https://travis-ci.com/your-user/your-repo/settings).

### Списывание
Домашние задания нужны **вам**, а не нам. За их невыполнение родителей в школу не вызовут.
Мимо нас проходят сотни решений и понимание того, прислал человек свой код или чужой,
происходит моментально. Это не влечёт за собой ничего, кроме как ухудшения отношения к нему,
как к потенциальному кандидату и будущему коллеге.

### Обратная связь
Сообщайте о найденных опечатках, ошибках, недочётах и пр.

Пожалуйста, не спешите писать о "неправильных тестах", если ваша программа
не проходит их. Сначала ещё раз внимательно прочитайте условие ДЗ
и проверьте написанный код. Спасибо! 
