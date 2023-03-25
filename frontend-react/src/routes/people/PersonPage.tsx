import { useParams } from "react-router-dom";
import { usePersonPageDataQuery } from "../../generated-types";
import ChatSection from "./ChatSection";
import { Avatar, Box, Container } from "@chakra-ui/react";

export default function PersonPage() {
  const params = useParams();
  const q = usePersonPageDataQuery({ variables: { id: params.id } });
  const pagePerson = q.data?.person;

  return (
    <Container>
      <Box>
        <Avatar size="lg" src={pagePerson?.avatarUrl} />
        <h1 className="px-2 text-6xl">{pagePerson?.name}</h1>
      </Box>
      {pagePerson && <ChatSection pagePerson={pagePerson} />}
    </Container>
  );
}
