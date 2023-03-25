--! Previous: sha1:1d7065cdd1017e95a319d17eb9ec3d8a717d325f
--! Hash: sha1:41199fb50f4c876838fb4319ffccafd330fcf7af

drop trigger if exists email_login_code on app_private.login_codes;
create trigger email_login_code after insert on app_private.login_codes for each row execute function app_public.trigger_job('email_login_code');
