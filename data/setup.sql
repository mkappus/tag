create table students (
	perm_id	integer primary key not null,
	grade integer,
	gu text,
	first_name text,
	last_name text
	);

create table student_course (
	perm_id integer,
	course_id integer
);

create table reservations (
	id integer primary key not null,
	title text,
	created integer,
	modified integer,
	start integer,
end integer,
creator_id integer,
course_id integer,
is_active tinyint
);

create table courses (
	id integer primary key not null,
	teacher_id integer,
	title text
);

create table teachers (
	id integer primary key not null,
	first_name text, 
	last_name text,
       email text,
	password text
);
