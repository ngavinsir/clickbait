alter table users rename column username to email;

alter table users add name text not null;

alter table users add is_male bool not null;

alter table users add age smallint not null;