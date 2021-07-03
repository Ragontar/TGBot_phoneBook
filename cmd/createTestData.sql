DROP TABLE  users;
DROP TABLE phonebook;


CREATE TABLE users (
                       user_id bigint,
                       phonebook integer[]
);

CREATE TABLE phonebook (
                           id SERIAL,
                           name varchar(200),
                           phone_number varchar(20)
);


INSERT INTO users VALUES (
                            100,
                            '{1, 2, 3}'
                         );

INSERT INTO users VALUES (
                             200,
                             '{2, 3, 5, 6}'
                         );

INSERT INTO phonebook VALUES (
                                1,
                                'Vitya Pitersky',
                              '11111'
                             );

INSERT INTO phonebook VALUES (
                                 2,
                                 'Jopa Pitersky',
                                 '22222'
                             );
INSERT INTO phonebook VALUES (
                                 3,
                                 'Ilyha Pisuyha',
                                 '33333'
                             );
INSERT INTO phonebook VALUES (
                                 4,
                                 'Vika the Woodcutter',
                                 '444444'
                             );
INSERT INTO phonebook VALUES (
                                 5,
                                 'Gleck Glechovich',
                                 '555555'
                             );
INSERT INTO phonebook VALUES (
                                 6,
                                 'George Perry Floyd Jr.',
                                 '66666'
                             );
INSERT INTO phonebook VALUES (
                                 7,
                                 'Old Hikkan',
                                 '777777'
                             );