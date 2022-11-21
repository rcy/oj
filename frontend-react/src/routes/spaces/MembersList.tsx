import { useContext } from "react"
import { PersonIdContext } from "../../contexts"
import { Scalars, useSpaceMembershipsBySpaceIdQuery } from "../../generated-types"

interface Props {
  spaceId: Scalars['UUID']
}

export default function MembersList({ spaceId }: Props) {
  const queryResult  = useSpaceMembershipsBySpaceIdQuery({ variables: { spaceId } })
  const personId = useContext(PersonIdContext)

  return (
    <div className="flex">
      {queryResult.data?.spaceMemberships?.edges?.map(x => (
        <div key={x.node.id}>
          <img alt="avatar" src={x.node.person?.avatarUrl} />
          <div>{x.node.person?.name}</div>
          <div>{x.node.person?.id === personId && '(me)'}</div>
        </div>
      ))}
    </div>
  )
}
