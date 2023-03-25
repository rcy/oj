import {
  Text,
  Heading,
  Box,
  Flex,
  CardBody,
  Card,
  Avatar,
  AvatarBadge,
} from "@chakra-ui/react";

import { Link } from "react-router-dom";

type PersonCardPrimitiveProps = {
  username: string;
  name?: string;
  avatarUrl?: string;
  title?: string;
  online: boolean;
  path: string;
};

export function PersonCardPrimitive(props: PersonCardPrimitiveProps) {
  return (
    <Link to={props.path}>
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

type PersonCardProps = {
  person: any;
};

export default function PersonCard(props: PersonCardProps) {
  const person = props.person;
  return (
    <PersonCardPrimitive
      username={person.username}
      name={person.name}
      avatarUrl={person.avatarUrl}
      title={person.title}
      online={false}
      path={`/people/${person.id}`}
    />
  );
}
