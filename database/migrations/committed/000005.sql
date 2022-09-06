--! Previous: sha1:9f1f2bca52283665cde704ca7bb9f38b317d35a9
--! Hash: sha1:a6ffffa2b07638496306989447a5fcaa99d782b5

-- enable rls on authentications, limit to showing only a users own records

alter table app_public.authentications enable row level security;

drop policy if exists select_own on app_public.authentications;
create policy select_own on app_public.authentications for select using (user_id = app_public.user_id());
