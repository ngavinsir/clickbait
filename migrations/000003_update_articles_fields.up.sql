alter table articles
	add date timestamptz not null;

alter table articles
	add source text not null;

alter table articles
	add url text not null;