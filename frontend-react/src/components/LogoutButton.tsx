import { logout } from "../util/logout";

export default function LogoutButton() {
  return <button onClick={logout}>logout</button>;
}
