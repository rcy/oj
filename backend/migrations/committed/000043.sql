--! Previous: sha1:6646c356527bebe216e04c209d1ab05c5f0d99e6
--! Hash: sha1:5b114d1a541ba64dec602d3feeec04313dc68ff0

-- Enter migration here
drop trigger if exists email_login_code on app_private.login_codes;
create trigger email_login_code after insert on app_private.login_codes for each row execute function app_public.trigger_job('email_login_code');
