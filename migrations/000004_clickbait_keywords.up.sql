create table clickbait_keywords
(
	id text not null
		constraint clickbait_keywords_pk
			primary key,
	label_id text not null
		constraint clickbait_keywords_labels_id_fk
			references labels
				on delete cascade,
	keyword text not null
);