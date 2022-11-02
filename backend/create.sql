create role appuser with password 'appuser' login;
create database app_development with owner appuser;
create database app_development_shadow with owner appuser;

create role visitor with password 'visitor' login;
grant all privileges on database app_development to visitor;
grant visitor to appuser;
