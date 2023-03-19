drop policy if exists select_own on app_public.families;
create policy select_all on app_public.families for select using (true);
