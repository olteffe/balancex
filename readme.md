Задание:
[тык сюда](https://github.com/avito-tech/autumn-2021-intern-assignment) или ./task.zip

Принятые решения: 
==================
1. Переводим с REST на gRPC


Временные заметки:
===================
1. мс биллинг ->: Добавить на баланс получателя.
2. мс услуги ->: Трата пользователем: Проверяем баланс(>= нужной суммы) получателя, фризим, списываем или отменяем
по запросу или ttl.
3. мс услуги ->: Перевод: Проверяем баланс отправителя(>= нужной суммы), фризим сумму отправителя, списываем у
отправителя, начисляем или отменяем в транзакции.