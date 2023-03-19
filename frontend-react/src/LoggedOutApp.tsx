import { Button, Center, HStack, VStack } from "@chakra-ui/react";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import PersonAuth from "./PersonAuth";

export default function LoggedOutApp() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/kidsauth/*" element={<PersonAuth />} />
        <Route path="*" element={<AuthButtons />} />
      </Routes>
    </BrowserRouter>
  );
}

function AuthButtons() {
  const here = encodeURIComponent(window.location.href);

  return (
    <Center h='100vh'>
      <VStack>
        <VStack>
          <img alt="octopus" width="300px" src="octopus1.png" />
          <img
            alt="octopus junior text"
            width="300px"
            src="octopus-junior-text.png"
          />
        </VStack>
        <Link to={`/kidsauth/login?from=${here}`}>
          <Button color="red">kids login</Button>
        </Link>
        <a href={`/auth/login?from=${here}`}>
          <Button color="blue">parents login</Button>
        </a>
      </VStack>
    </Center>
  );
}
