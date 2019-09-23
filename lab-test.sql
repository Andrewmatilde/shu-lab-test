CREATE EXTENSION pgcrypto;
create table student
(
    id       char(8) not null,
    password text    not null
);

create unique index student_id_uindex
    on student (id);

alter table student
    add constraint student_pk
        primary key (id);