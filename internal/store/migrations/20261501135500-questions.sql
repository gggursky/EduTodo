-- +migrate Up

INSERT INTO questions (code,name,description)
VALUES ('code1','вопрос №1', 'Что является обязательным элементом при монтаже панелей ЗИПС на стену или потолок?');

INSERT INTO questions (code,name,description)
VALUES ('code2','вопрос №2', 'Какая система ЗИПС обеспечивает наибольший индекс дополнительной звукоизоляции ΔRw?');

INSERT INTO questions (code,name,description)
VALUES ('code3','вопрос №1', 'Что является основной причиной того, что многослойные конструкции обеспечивают более ' ||
                             'высокую звукоизоляцию по сравнению с однослойными?');

INSERT INTO questions (code,name,description)
VALUES ('code4','вопрос №2', 'Что в наибольшей степени улучшает защиту от ударного шума в конструкции «плавающего пола»?');

INSERT INTO questions (code,name,description)
VALUES ('code5','вопрос №1', 'Что произойдёт, если при монтаже звукоизоляционной конструкции не использовать вибропрокладки?');

INSERT INTO questions (code,name,description)
VALUES ('code6','вопрос №2', 'Какой материал нельзя использовать в качестве поглощающего слоя внутри каркасной звукоизоляционной системы?');