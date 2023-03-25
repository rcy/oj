import formData from 'form-data';
import Mailgun from 'mailgun.js';

export async function email_login_code(payload, helpers) {
  const mailgun = new Mailgun(formData);
  const mg = mailgun.client({ username: 'api', key: process.env.MAILGUN_API_KEY });

  const { rows: [{ user_id, code }] } = await helpers.query(`
select users.id as user_id, code
from app_private.login_codes
join app_public.managed_people on managed_people.person_id = login_codes.person_id
join app_public.users on users.id = managed_people.user_id
where login_codes.id = $1
  `, [payload.id])

  // lookup the email for the user from authentications.
  // we should have a users table which is populated when the user is created and not dig around in the authentications records
  const { rows: [{ email }] } = await helpers.query(`
select details->'_json'->>'email' as email from app_public.authentications where user_id = $1
  `, [user_id])

  console.log({ user_id, code, email })

  // send the email
  await mg.messages
    .create('mg.ryanyeske.com', {
      from: "Octopus Jr. <mailgun@mg.ryanyeske.com>",
      to: [email],
      subject: `Here's your kid's login code ${code}`,
      text: `Your kid's login code is ${code}`,
      html: `Your kid's login code is <h1>${code}</h1>`
    })
};
