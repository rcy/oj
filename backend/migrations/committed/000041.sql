--! Previous: sha1:f3574ed6b57b0330ece73351add6307d12451234
--! Hash: sha1:1d7065cdd1017e95a319d17eb9ec3d8a717d325f

drop policy if exists select_all on app_public.families;
drop policy if exists select_own on app_public.families;
create policy select_all on app_public.families for select using (true);
