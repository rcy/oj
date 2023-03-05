--! Previous: sha1:c130275f10cfd2d66e3b95c4a377bd433efd2ddd
--! Hash: sha1:4c667aa3df145d18dc5384d62260f2d972b972ae

-- Enter migration here
drop table if exists app_public.notifications;
create table app_public.notifications (
  id uuid default gen_random_uuid() not null,
  created_at timestamp with time zone default now() not null,
  post_id uuid references app_public.posts not null,
  membership_id uuid references app_public.space_memberships not null
);

grant select on table app_public.notifications to visitor;

drop trigger if exists create_post_notifications on app_public.posts;
create trigger create_post_notifications
  after insert on app_public.posts
  for each row
  execute procedure trigger_job('create_post_notifications');
