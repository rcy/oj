import { Flex, Box, Text } from "@chakra-ui/react";
import { Link } from "react-router-dom";

export default function ParentAdminHeader() {
  return (
    <Flex as="nav" alignItems="center" p="2" boxShadow="md">
      <Box>
        <Text fontSize="xl" fontWeight="bold">
          <Link to="/parent">Octopus Jr. Parent</Link>
        </Text>
      </Box>
    </Flex>
  );
}
