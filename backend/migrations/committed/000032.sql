--! Previous: sha1:4c667aa3df145d18dc5384d62260f2d972b972ae
--! Hash: sha1:ead6262d24bda06e44c1dddaad3cb0a23cc48980

-- Enter migration here
drop trigger if exists _500_notify on app_public.posts;
drop function if exists app_public.notify_space_post_created;

create function app_public.notify_space_post_created() returns trigger as $$
begin
  perform pg_notify('graphql:spaceposts:' || new.space_id, json_build_object('event', 'postCreated', 'subject', new.id)::text);
  return new;
end
$$ language plpgsql volatile set search_path from current;

create trigger _500_notify
  after insert on app_public.posts
  for each row
  execute procedure app_public.notify_space_post_created();
