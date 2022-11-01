--! Previous: sha1:40a038759ef95a4c2e46a42e6454041f221777c8
--! Hash: sha1:ab3bb59208d99fc0c0043abf0664c64a0f32a0b6

-- add timestamps to family memberships

alter table app_public.family_memberships
add column if not exists created_at timestamptz not null default now();
