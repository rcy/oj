insert into users(username, email, is_parent) values('p1', 'p1@example.com', 1);

insert into users(username, email) values('p1c1', 'p1c1@example.com');
insert into kids_parents(kid_id, parent_id) values(
  (select id from users where username = 'p1c1'),
  (select id from users where username='p1')
);
insert into friends(a_id, b_id, b_role) values(
  (select id from users where username = 'p1c1'),
  (select id from users where username = 'p1'),
  'parent'
);
insert into friends(a_id, b_id, b_role) values(
  (select id from users where username = 'p1'),
  (select id from users where username = 'p1c1'),
  'child'
);

insert into users(username, email) values('p1c2', 'p1c2@example.com');
insert into kids_parents(kid_id, parent_id) values(
  (select id from users where username = 'p1c2'),
  (select id from users where username='p1')
);
insert into friends(a_id, b_id, b_role) values(
  (select id from users where username = 'p1c2'),
  (select id from users where username = 'p1'),
  'parent'
);
insert into friends(a_id, b_id, b_role) values(
  (select id from users where username = 'p1'),
  (select id from users where username = 'p1c2'),
  'child'
);


insert into users(username, email, is_parent) values('p2', 'p2@example.com', 1);

insert into users(username, email) values('p2c1', 'p2c1@example.com');
insert into kids_parents(kid_id, parent_id) values(
  (select id from users where username = 'p2c1'),
  (select id from users where username='p2')
);
insert into friends(a_id, b_id, b_role) values(
  (select id from users where username = 'p2c1'),
  (select id from users where username = 'p2'),
  'parent'
);
insert into friends(a_id, b_id, b_role) values(
  (select id from users where username = 'p2'),
  (select id from users where username = 'p2c1'),
  'child'
);

insert into users(username, email) values('p2c2', 'p2c2@example.com');
insert into kids_parents(kid_id, parent_id) values(
  (select id from users where username = 'p2c2'),
  (select id from users where username='p2')
);
insert into friends(a_id, b_id, b_role) values(
  (select id from users where username = 'p2c2'),
  (select id from users where username = 'p2'),
  'parent'
);
insert into friends(a_id, b_id, b_role) values(
  (select id from users where username = 'p2'),
  (select id from users where username = 'p2c2'),
  'child'
);

