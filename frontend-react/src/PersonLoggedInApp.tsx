import LogoutButton from "./components/LogoutButton";
import { Text, Heading, Box, Flex, Spacer, Container, CardBody, Card, Avatar } from '@chakra-ui/react'
import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import MemberSetProfilePicture from "./routes/member/MemberSetProfilePicture";
import { useCurrentPersonQuery } from "./generated-types";

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
  const personQuery = useCurrentPersonQuery({
    fetchPolicy: "network-only",
  });

  if (personQuery.loading) {
    return <div></div>
  }

  return (
    <div>
      <PersonCard
        name={personQuery.data?.currentPerson?.name}
        avatarUrl={personQuery.data?.currentPerson?.avatarUrl}
        username={personQuery.data?.currentPerson?.username || ''}
        title="me"
      />
    </div>
  )
}

type PersonCardProps = {
  username: string
  name?: string
  avatarUrl?: string
  title?: string
}

function PersonCard(props: PersonCardProps) {
  return (
    <Link to="/pic">
      <Card maxW='xs'>
        <CardBody>
          <Flex>
            <Flex flex='1' gap='4' alignItems='center' flexWrap='wrap'>
              <Avatar
                size='lg'
                name={props.name}
                src={props.avatarUrl}
              />
              <Box>
                <Heading size='lg'>{props.username}</Heading>
                <Text>{props.title}</Text>
              </Box>
            </Flex>
          </Flex>
        </CardBody>
      </Card>
    </Link>
  )
}
