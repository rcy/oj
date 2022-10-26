alter table app_public.people add column if not exists avatar_url text not null default 'https://www.gravatar.com/avatar/DEFAULT?f=y&d=mp';
