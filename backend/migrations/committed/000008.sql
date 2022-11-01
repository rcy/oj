--! Previous: sha1:b0735dca145fdb4c63a68b4e6fc02cd1d42aa830
--! Hash: sha1:be1904ec9b1ce61f5e50bedc38b73adeed190ed0

revoke all privileges on table app_public.authentications from visitor;
grant select on table app_public.authentications to visitor;
