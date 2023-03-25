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
import PersonPage from "./routes/people/PersonPage";
import { PersonCardPrimitive } from "./components/PersonCard";

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
        <Route path="/people/:id" element={<PersonPage />} />
      </Routes>
    </Box>
  );
}

function Main() {
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
          <PersonCardPrimitive
            key={m.node.id}
            name={m.node.person?.name}
            avatarUrl={m.node.person?.avatarUrl}
            username={m.node.person?.username || ""}
            title={m.node.person?.name}
            online={false}
            path={`/people/${m.node.person?.id}`}
          />
        ))}
      </SimpleGrid>
    </VStack >
  );
}
