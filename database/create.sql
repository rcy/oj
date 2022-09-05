create role appuser with password 'appuser' login;
create database app_development with owner appuser;
create database app_development_shadow with owner appuser;
