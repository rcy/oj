--! Previous: sha1:ab3bb59208d99fc0c0043abf0664c64a0f32a0b6
--! Hash: sha1:f6f2b2e19faa254887256d90583ad87b3b03fe02

alter table app_public.people add column if not exists avatar_url text not null default 'https://www.gravatar.com/avatar/DEFAULT?f=y&d=mp';
