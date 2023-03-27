--! Previous: sha1:1d7065cdd1017e95a319d17eb9ec3d8a717d325f
--! Hash: sha1:6646c356527bebe216e04c209d1ab05c5f0d99e6

-- trigger_job got created on the public schema on prod in 000030.sql, untangle that here

drop trigger if exists create_post_notifications on app_public.posts;

drop function if exists public.trigger_job;
drop function if exists app_public.trigger_job;
create function app_public.trigger_job() returns trigger as $$
begin
  perform graphile_worker.add_job(tg_argv[0], json_build_object(
    'schema', tg_table_schema,
    'table', tg_table_name,
    'op', tg_op,
    'id', (case when tg_op = 'delete' then old.id else new.id end)
  ));
  return new;
end;
$$ language plpgsql volatile;


create trigger create_post_notifications after insert on app_public.posts for each row execute function app_public.trigger_job('create_post_notifications');
