import LogoutButton from "./components/LogoutButton";
import { Heading, Box, Flex, Spacer, VStack, Container } from '@chakra-ui/react'
import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import Chat from "./components/Chat";
import MemberSetProfilePicture from "./routes/member/MemberSetProfilePicture";

export default function PersonLoggedInApp() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="*" element={<PersonLoggedInAppInner />} />
      </Routes>
    </BrowserRouter>
  );
}

function PersonLoggedInAppInner() {
  return (
    <Box minWidth='max-content'>
      <Box minWidth='max-content' background="purple.300">
        <Container minWidth={1000}>
          <Flex alignItems='center' py='2' mb='1em'>
            <Box>
              <Heading size="md">
                <Link to="/">Octopus Jr.</Link>
              </Heading>
            </Box>
            <Spacer />
            <Box>
              <LogoutButton />
            </Box>
          </Flex>
        </Container >
      </Box>
      <Container minWidth={1000}>
        <Routes>
          <Route path="/" element={<Main />} />
          <Route path="/pic" element={<MemberSetProfilePicture />} />
        </Routes>
      </Container>
    </Box >
  )
}

function Main() {
  return <Link to="/pic">Change Profile Picture</Link>
}
