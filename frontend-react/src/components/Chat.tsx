import { useEffect } from "react";
import {
  Post,
  SpacePostsAddedDocument,
  useCurrentPersonQuery,
  usePostMessageMutation,
  useSpaceMembershipByPersonIdAndSpaceIdQuery,
  useSpacePostsQuery,
} from "../generated-types";
import ChatInput from "./ChatInput";
import beep from "../util/beep";
import { Flex, Text, Card, Avatar, Spacer, Box } from "@chakra-ui/react";

interface Props {
  spaceId: string;
}

export default function Chat({ spaceId }: Props) {
  const spacePostsQueryResult = useSpacePostsQuery({
    variables: { spaceId, limit: 10 },
  });
  const [postMessageMutation] = usePostMessageMutation();
  const personQuery = useCurrentPersonQuery();
  const personId = personQuery.data?.currentPerson?.id;
  const membershipQueryResult = useSpaceMembershipByPersonIdAndSpaceIdQuery({
    variables: { spaceId, personId },
  });

  const membershipId =
    membershipQueryResult.data?.spaceMembershipByPersonIdAndSpaceId?.id;

  const person = personQuery.data?.currentPerson

  useEffect(() => {
    // subscribeToMore returns its unsubscribe function
    return spacePostsQueryResult.subscribeToMore({
      document: SpacePostsAddedDocument,
      variables: { spaceId },
      updateQuery: (prev, { subscriptionData }) => {
        const p: any = subscriptionData.data.posts;

        if (prev.posts) {
          // this could be handled by notifications in the future
          if (p.post.membership.id !== membershipId) {
            beep();
          }

          return {
            ...prev,
            posts: {
              ...prev.posts,
              nodes: [...prev.posts.nodes, p.post].slice(-10),
            },
          };
        }
        return prev;
      },
    });
  }, [spaceId, membershipId, spacePostsQueryResult]);

  const handleSubmit = async (text: string) => {
    await postMessageMutation({
      variables: {
        membershipId:
          membershipQueryResult.data?.spaceMembershipByPersonIdAndSpaceId?.id,
        body: text,
      },
    });

    // refetch all the messages (switch to subscription)
    //spacePostsQueryResult.refetch();

    return true;
  };

  return (
    <Flex direction="column">
      <Box height="50vh">
        {spacePostsQueryResult.data?.posts?.nodes.map((post) => (
          <ChatPost key={post.id} post={post} />
        ))}
      </Box>
      <Spacer/>
      <ChatInput onSubmit={handleSubmit} person={person} />
    </Flex>
  );
}

function ChatPost(props: any) {
  const post: Post = props.post

  const person = post.membership?.person

  const date = new Date(post.createdAt)
  const formattedTime = date.toLocaleString('en-US', {
    hour: "numeric",
    minute: "numeric",
    hour12: true
  })

  return (
    <Card>
      <Flex gap="1">
        <Avatar size="xs" src={person?.avatarUrl} name={person?.name} />
        <Text fontWeight="bold">{person?.name}: </Text>
        <Text flex="1">{post.body}</Text>
        <Text color="gray.500" flex="1" maxWidth="5em" fontSize="xs" align="right">{formattedTime}</Text>
      </Flex>
    </Card>
  )
}
