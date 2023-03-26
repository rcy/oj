import {
  Avatar,
  Box,
  Container,
  Button,
  Heading,
  SimpleGrid,
} from "@chakra-ui/react";
import md5 from "md5";
import { useEffect, useState } from "react";
import {
  useCurrentPersonQuery,
  useSetPersonAvatarMutation,
} from "../../generated-types";
import LoggedOutApp from "../../LoggedOutApp";
import LogoutButton from "../../components/LogoutButton";

// return a hash based on a given string and index
function getDailyHash(str, index) {
  return md5(`${str}${index}`);
}

const types = ["monsterid", "wavatar", "robohash", "identicon", "retro"];

export default function MemberSetProfilePicture() {
  const currentPersonQuery = useCurrentPersonQuery();
  const personId = currentPersonQuery?.data?.currentPerson.id;
  const [dtype, selectType] = useState(types[0]);
  const [mutation] = useSetPersonAvatarMutation();
  const [hashes, setHashes] = useState([]);

  async function handleSelect(avatarUrl) {
    await mutation({ variables: { personId, avatarUrl } });
    currentPersonQuery.refetch();
  }

  function shuffle() {
    setHashes([...Array(32).keys()].map((k) => getDailyHash(personId, k)));
  }

  useEffect(shuffle, [personId]);

  return (
    <Container>
      <Box textAlign="right">
        <LogoutButton />
      </Box>

      <Heading as="h1">Change your profile picture</Heading>

      <Box py="5">
        {types.map((d) => (
          <Button
            onClick={() => selectType(d)}
            key={d}
            colorScheme="green"
            m={1}
          >
            {d}
          </Button>
        ))}
      </Box>

      <SimpleGrid minChildWidth="100px" spacing="5px">
        {hashes.map((h) => (
          <div key={h}>
            <GravImage h={h} d={dtype} s={80} onSelect={handleSelect} />
          </div>
        ))}
      </SimpleGrid>
    </Container>
  );
}

function GravImage({ h, d, s, onSelect }) {
  const url = `https://www.gravatar.com/avatar/${h}?f=y&d=${d}`;

  function handleClick(ev) {
    ev.preventDefault();
    return onSelect(url);
  }

  return (
    <Avatar
      alt="avatar"
      size="xl"
      key={h}
      src={`${url}&s=${s}`}
      onClick={handleClick}
    />
  );
}
