--! Previous: sha1:5b114d1a541ba64dec602d3feeec04313dc68ff0
--! Hash: sha1:9aef864ac32193e3e8f5266b8dabc22c0e9194fe

-- our avatar component can handle generating a good default for us based on the name
alter table app_public.people alter column avatar_url set default '';
