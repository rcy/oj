import {
  VStack,
  Text,
  Heading,
  Box,
  Flex,
  CardBody,
  Card,
  Avatar,
  SimpleGrid,
  AvatarBadge,
} from "@chakra-ui/react";
import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import MemberSetProfilePicture from "./routes/member/MemberSetProfilePicture";
import { useCurrentPersonFamilyMembershipQuery, useCurrentPersonQuery } from "./generated-types";
import NavBar from "./components/NavBar";

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
    <Box
      background={"red.100"}
      minH="100vh"
      bgGradient="linear(to-r, orange.300, purple.300)"
    >
      <NavBar />
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/profile" element={<MemberSetProfilePicture />} />
      </Routes>
    </Box>
  );
}

function Main() {
  const personQuery = useCurrentPersonQuery({
    fetchPolicy: "network-only",
  });

  const membershipQuery = useCurrentPersonFamilyMembershipQuery()

  const currentPerson = membershipQuery.data?.currentPerson
  const memberships = currentPerson?.familyMembership?.family?.familyMemberships.edges
  const filteredMemberships = memberships?.filter(edge => edge.node.person?.id !== currentPerson?.id)

    console.log({ memberships })

  return (
    <VStack>
      <Heading>My Family</Heading>
      <SimpleGrid columns={1} gap="20px">
        {filteredMemberships?.map(m => (
          <PersonCard
            key={m.node.id}
            name={m.node.person?.name}
            avatarUrl={m.node.person?.avatarUrl}
            username={m.node.person?.username || ""}
            title="me"
            online
          />
        ))}
      </SimpleGrid>
    </VStack >
  );
}

type PersonCardProps = {
  username: string;
  name?: string;
  avatarUrl?: string;
  title?: string;
  online: boolean;
};

function PersonCard(props: PersonCardProps) {
  return (
    <Link to="/profile">
      <Card w="xs">
        <CardBody>
          <Flex>
            <Flex flex="1" gap="4" alignItems="center" flexWrap="wrap">
              <Avatar size="lg" name={props.name} src={props.avatarUrl}>
                {props.online && <AvatarBadge boxSize="1em" bg="green.500" />}
              </Avatar>
              <Box>
                <Heading size="lg">{props.username}</Heading>
                <Text>{props.title}</Text>
              </Box>
            </Flex>
          </Flex>
        </CardBody>
      </Card>
    </Link>
  );
}
