--! Previous: sha1:dfea05713d55236117820fd5b35321980d5a533f
--! Hash: sha1:c130275f10cfd2d66e3b95c4a377bd433efd2ddd

-- Enter migration here

drop function if exists trigger_job;
CREATE FUNCTION trigger_job() RETURNS trigger AS $$
BEGIN
  PERFORM graphile_worker.add_job(TG_ARGV[0], json_build_object(
    'schema', TG_TABLE_SCHEMA,
    'table', TG_TABLE_NAME,
    'op', TG_OP,
    'id', (CASE WHEN TG_OP = 'DELETE' THEN OLD.id ELSE NEW.id END)
  ));
  RETURN NEW;
END;
$$ LANGUAGE plpgsql VOLATILE;
