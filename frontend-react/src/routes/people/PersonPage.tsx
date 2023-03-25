import { useParams } from "react-router-dom";
import { usePersonPageDataQuery } from "../../generated-types";
import ChatSection from "./ChatSection";
import { Container, Spinner } from "@chakra-ui/react";
import PersonCard from "../../components/PersonCard";

export default function PersonPage() {
  const params = useParams();
  const q = usePersonPageDataQuery({ variables: { id: params.id } });
  const pagePerson = q.data?.person;

  if (q.loading) {
    return <Spinner />
  }

  return (
    <Container>
      <PersonCard person={pagePerson} />
      {pagePerson && <ChatSection pagePerson={pagePerson} />}
    </Container>
  );
}
