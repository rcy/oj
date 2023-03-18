import { useContext, useEffect } from "react";
import { PersonIdContext } from "../contexts";
import {
  SpacePostsAddedDocument,
  usePostMessageMutation,
  useSpaceMembershipByPersonIdAndSpaceIdQuery,
  useSpacePostsQuery,
} from "../generated-types";
import ChatInput from "./ChatInput";
import beep from "../util/beep";

interface Props {
  spaceId: string;
}

export default function Chat({ spaceId }: Props) {
  const spacePostsQueryResult = useSpacePostsQuery({
    variables: { spaceId, limit: 10 },
  });
  const [postMessageMutation] = usePostMessageMutation();
  const personId = useContext(PersonIdContext);
  const membershipQueryResult = useSpaceMembershipByPersonIdAndSpaceIdQuery({
    variables: { spaceId, personId },
  });

  const membershipId =
    membershipQueryResult.data?.spaceMembershipByPersonIdAndSpaceId?.id;

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
    <div>
      {spacePostsQueryResult.data?.posts?.nodes.map((post) => (
        <div key={post.id} className="flex">
          <img
            alt="avatar"
            width="32"
            src={post.membership?.person?.avatarUrl}
          />
          <b>{post.membership?.person?.name}: </b>
          <div>{post.body}</div>
        </div>
      ))}

      <ChatInput onSubmit={handleSubmit} />
    </div>
  );
}
