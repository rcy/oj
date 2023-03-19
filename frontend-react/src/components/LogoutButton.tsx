import { Button } from "@chakra-ui/react";
import { logout } from "../util/logout";

export default function LogoutButton() {
  function click() {
    if (window.confirm("are you sure you want to logout?")) {
      logout();
    }
  }
  return (
    <Button colorScheme="blue" onClick={click}>
      logout
    </Button>
  );
}
