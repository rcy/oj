--! Previous: sha1:602c804a14e22121cc941a169904a19b40effa19
--! Hash: sha1:40a038759ef95a4c2e46a42e6454041f221777c8

delete from app_public.posts;

-- add space_id to posts
alter table app_public.posts
drop column if exists space_id;

alter table app_public.posts
add column space_id uuid references app_public.spaces not null;

-- create posts that sets the space_id from the membership.space_id
create or replace function app_public.post_message(space_membership_id uuid, body text)
returns app_public.posts
language plpgsql strict security definer
as $$
declare
        v_space_id uuid;
        v_result app_public.posts;
begin
        select space_id
        into v_space_id
        from app_public.space_memberships
        where id = space_membership_id;

        insert
                into app_public.posts(membership_id, space_id, body)
                values(space_membership_id, v_space_id, body)
                returning * into v_result;

        return v_result;
end;
$$;

-- disallow inserting directly into posts
revoke all on table app_public.posts from visitor;
grant select on table app_public.posts to visitor;
