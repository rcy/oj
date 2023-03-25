import {
  Heading,
  Box,
  Flex,
  Spacer,
  Container,
  Avatar,
  Text,
  HStack,
} from "@chakra-ui/react";
import { Link } from "react-router-dom";
import { useCurrentPersonQuery } from "../generated-types";

export default function NavBar() {
  const query = useCurrentPersonQuery();
  console.log({ query });

  const person = query.data?.currentPerson;

  return (
    <Box
      minWidth="max-content"
      bgGradient="linear(to-r, orange.200, purple.200)"
    >
      <Container>
        <Flex alignItems="center" py="2" mb="1em">
          <Heading size="lg">
            <Link to="/">🐙</Link>
            <Link to="/"> Octopus Jr.</Link>
          </Heading>
          <Spacer />
          <Box alignContent="center">
            <Link to="/profile">
              <HStack>
                <Text>{person?.username}</Text>
                <Avatar src={person?.avatarUrl} />
              </HStack>
            </Link>
          </Box>
        </Flex>
      </Container>
    </Box>
  );
}
