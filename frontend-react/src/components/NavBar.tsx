import {
  Center,
  Text,
  Heading,
  Box,
  Flex,
  Spacer,
  Container,
  CardBody,
  Card,
  Avatar,
  SimpleGrid,
  AvatarBadge,
} from "@chakra-ui/react";
import { BrowserRouter, Link, Route, Routes } from "react-router-dom";

export default function NavBar() {
  return (
    <Box
      minWidth="max-content"
      bgGradient="linear(to-r, orange.200, purple.200)"
    >
      <Container>
        <Flex alignItems="center" py="2" mb="1em">
          <Heading size="xl">
            <Link to="/">🐙</Link>
          </Heading>
          <Spacer />
          <Box>
            <Heading size="md">
              <Link to="/">Octopus Junior</Link>
            </Heading>
          </Box>
          <Spacer />
          <Box>
            <Link to="/profile">
              <Avatar />
            </Link>
          </Box>
        </Flex>
      </Container>
    </Box>
  );
}
