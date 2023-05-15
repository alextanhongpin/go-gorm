-- show create table user;

create table user (
	id bigint(20) unsigned primary key auto_increment,
	name varchar(255) not null,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp on update current_timestamp,
	unique `user_name_uk`(name)
);

insert into user (name) values ('john');
insert into user (name) values
('john')
on duplicate key update
	updated_at = current_timestamp();
