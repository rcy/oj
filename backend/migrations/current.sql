-- Enter migration here
drop trigger if exists _500_gql_update on app_public.users;

-- from https://www.graphile.org/postgraphile/subscriptions
drop function if exists app_public.graphql_subscription;
create function app_public.graphql_subscription() returns trigger as $$
declare
  v_process_new bool = (TG_OP = 'INSERT' OR TG_OP = 'UPDATE');
  v_process_old bool = (TG_OP = 'UPDATE' OR TG_OP = 'DELETE');
  v_event text = TG_ARGV[0];
  v_topic_template text = TG_ARGV[1];
  v_attribute text = TG_ARGV[2];
  v_record record;
  v_sub text;
  v_topic text;
  v_i int = 0;
  v_last_topic text;
begin
  -- On UPDATE sometimes topic may be changed for NEW record,
  -- so we need notify to both topics NEW and OLD.
  for v_i in 0..1 loop
    if (v_i = 0) and v_process_new is true then
      v_record = new;
    elsif (v_i = 1) and v_process_old is true then
      v_record = old;
    else
      continue;
    end if;
     if v_attribute is not null then
      execute 'select $1.' || quote_ident(v_attribute)
        using v_record
        into v_sub;
    end if;
    if v_sub is not null then
      v_topic = replace(v_topic_template, '$1', v_sub);
    else
      v_topic = v_topic_template;
    end if;
    if v_topic is distinct from v_last_topic then
      -- This if statement prevents us from triggering the same notification twice
      v_last_topic = v_topic;
      perform pg_notify(v_topic, json_build_object(
        'event', v_event,
        'subject', v_sub
      )::text);
    end if;
  end loop;
  return v_record;
end;
$$ language plpgsql volatile set search_path from current;




-- delete this later
CREATE TRIGGER _500_gql_update
  AFTER UPDATE ON app_public.users
  FOR EACH ROW
  EXECUTE PROCEDURE app_public.graphql_subscription(
    'userChanged', -- the "event" string, useful for the client to know what happened
    'graphql:user:$1', -- the "topic" the event will be published to, as a template
    'id' -- If specified, `$1` above will be replaced with NEW.id or OLD.id from the trigger.
  );


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
