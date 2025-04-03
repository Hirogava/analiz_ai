DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM "public"."categories") THEN
        INSERT INTO "public"."categories" (name) VALUES ('Проблемы с авторизацией');
        INSERT INTO "public"."categories" (name) VALUES ('Проблемы с оформлением заказа');
        INSERT INTO "public"."categories" (name) VALUES ('Проблемы с поиском');
        INSERT INTO "public"."categories" (name) VALUES ('Проблемы с отображением страниц');
        INSERT INTO "public"."categories" (name) VALUES ('Технические вопросы');
        INSERT INTO "public"."categories" (name) VALUES ('Другие вопросы');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM "public"."questions") THEN
        INSERT INTO "public"."questions" (category_id, question) VALUES (1, 'Я не могу войти в свой аккаунт.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (1, 'Забыл пароль.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (1, 'Не могу зарегистрироваться.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (1, 'Двойная регистрация одного аккаунта.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (1, 'Проблемы с входом через социальные сети.');

        INSERT INTO "public"."questions" (category_id, question) VALUES (2, 'Корзина пустая после добавления товара.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (2, 'Ошибка при оформлении заказа.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (2, 'Проблемы с выбором доставки.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (2, 'Неверная сумма заказа.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (2, 'Не могу оплатить заказ.');

        INSERT INTO "public"."questions" (category_id, question) VALUES (3, 'Поиск не работает.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (3, 'Неправильные результаты поиска.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (3, 'Поиск медленный.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (3, 'Не могу найти нужный товар.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (3, 'Поиск не учитывает синонимы.');

        INSERT INTO "public"."questions" (category_id, question) VALUES (4, 'Страница не загружается.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (4, 'Страница отображается некорректно.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (4, 'Отсутствуют изображения.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (4, 'Проблемы с мобильной версией сайта.');
        INSERT INTO "public"."questions" (category_id, question) VALUES (4, 'Страница долго загружается.');

        INSERT INTO "public"."questions" (category_id, question) VALUES (1, 'Как сбросить пароль?');
        INSERT INTO "public"."questions" (category_id, question) VALUES (5, 'Как связаться с техподдержкой?');
        INSERT INTO "public"."questions" (category_id, question) VALUES (2, 'Как вернуть товар?');
        INSERT INTO "public"."questions" (category_id, question) VALUES (5, 'Как отменить заказ?');
        INSERT INTO "public"."questions" (category_id, question) VALUES (5, 'Как изменить информацию в профиле?');
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM "public"."answers") THEN
        INSERT INTO "public"."answers" (question_id, answer) VALUES (1, 'Проверьте правильность введенных данных и попробуйте снова.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (2, 'Перейдите на страницу восстановления пароля.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (3, 'Убедитесь, что все поля заполнены корректно.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (4, 'Проверьте почту для подтверждения регистрации.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (5, 'Попробуйте использовать другой метод входа.');

        INSERT INTO "public"."answers" (question_id, answer) VALUES (6, 'Обновите страницу или перезагрузите браузер.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (7, 'Проверьте корректность введенных данных и попробуйте снова.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (8, 'Выберите доступные варианты доставки.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (9, 'Проверьте корзину и пересчитайте сумму.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (10, 'Проверьте подключение к интернету и выберите другой метод оплаты.');

        INSERT INTO "public"."answers" (question_id, answer) VALUES (11, 'Проверьте подключение к интернету и попробуйте снова.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (12, 'Используйте точные ключевые слова.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (13, 'Попробуйте уменьшить количество ключевых слов.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (14, 'Проверьте правильность написания и попробуйте снова.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (15, 'Используйте синонимы вручную.');

        INSERT INTO "public"."answers" (question_id, answer) VALUES (16, 'Проверьте подключение к интернету и обновите страницу.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (17, 'Попробуйте очистить кэш браузера.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (18, 'Проверьте подключение к интернету и обновите страницу.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (19, 'Попробуйте использовать настольную версию сайта.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (20, 'Попробуйте очистить кэш браузера и обновить страницу.');

        INSERT INTO "public"."answers" (question_id, answer) VALUES (21, 'Перейдите на страницу восстановления пароля.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (22, 'Позвоните по телефону или напишите на электронную почту.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (23, 'Свяжитесь с службой поддержки для возврата товара.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (24, 'Свяжитесь с службой поддержки для отмены заказа.');
        INSERT INTO "public"."answers" (question_id, answer) VALUES (25, 'Перейдите в личный кабинет и отредактируйте данные.');
    END IF;
END $$;