create table if not exists users
(
	id varchar(50) not null
		constraint users_pk
			primary key,
	username varchar(50) not null,
	password varchar(100) not null
);

create unique index if not exists users_username_uindex
	on users (username);

create table if not exists articles
(
	id varchar(50) not null
		constraint headlines_pk
			primary key,
	headline text not null,
	content text not null
);

create table if not exists labels
(
	id varchar(50) not null
		constraint labels_pk
			primary key,
	user_id varchar(50) not null
		constraint labels_users_id_fk
			references users
				on delete cascade,
	article_id varchar(50) not null
		constraint labels_articles_id_fk
			references articles
				on delete cascade,
	value text not null,
	type varchar(20) not null,
	updated_at timestamp with time zone not null,
	created_at timestamp with time zone not null
);
