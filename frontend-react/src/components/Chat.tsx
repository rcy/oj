import { useContext } from "react"
import { PersonIdContext } from "../contexts"
import { usePostMessageMutation, useSpaceMembershipByPersonIdAndSpaceIdQuery, useSpacePostsQuery } from "../generated-types"
import ChatInput from './ChatInput'

interface Props {
  spaceId: string,
}

export default function Chat({ spaceId }: Props) {
  const spacePostsQueryResult = useSpacePostsQuery({ variables: { spaceId } });
  const [postMessageMutation] = usePostMessageMutation();
  const personId = useContext(PersonIdContext)
  const membershipQueryResult = useSpaceMembershipByPersonIdAndSpaceIdQuery({ variables: { spaceId, personId } })

  const handleSubmit = async (text: string) => {
    const result = await postMessageMutation({
      variables: {
        membershipId: membershipQueryResult.data?.spaceMembershipByPersonIdAndSpaceId?.id,
        body: text
      }
    })
    console.log({ result })

    // refetch all the messages (switch to subscription)
    spacePostsQueryResult.refetch()
    
    return true
  }

  return (
    <div>
      <h1>{spaceId}</h1>
      <hr />

      {spacePostsQueryResult.data?.posts?.edges.map(({ node: post }) => (
        <div key={post.id}>
          <b>{post.membership?.person?.name}</b>: {post.body}
        </div>
      ))}

      <ChatInput onSubmit={handleSubmit} />
    </div>
  )
}
