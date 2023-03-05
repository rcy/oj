export async function create_post_notifications(payload, helpers) {
  // get the post
  const { rows: [ post ] } = await helpers.query(`select * from app_public.posts where id = $1`, [payload.id])

  // get the memberships for the users in the space where the post was made
  const { rows: memberships } = await helpers.query(`select * from app_public.space_memberships where space_id = $1`, [post.space_id])

  // create notifications for each membership (not including sender)
  for (let membership of memberships) {
    if (membership.id !== post.membership_id) {
      await helpers.query(`insert into app_public.notifications(post_id, membership_id) values ($1, $2)`, [post.id, membership.id])
    }
  }
};
