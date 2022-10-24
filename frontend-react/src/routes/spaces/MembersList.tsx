import { useContext } from "react"
import Debug from "../../components/Debug"
import { PersonIdContext } from "../../contexts"
import { Scalars, useSpaceMembershipsBySpaceIdQuery } from "../../generated-types"

interface Props {
  spaceId: Scalars['UUID']
}

export default function MembersList({ spaceId }: Props) {
  const queryResult  = useSpaceMembershipsBySpaceIdQuery({ variables: { spaceId } })
  const personId = useContext(PersonIdContext)

  return (
    <div>
      {queryResult.data?.spaceMemberships?.edges?.map(x => (
        <div key={x.node.id}>
          {x.node.person?.name}
          {x.node.person?.id === personId && '(me)'}
        </div>
      ))}
    </div>
  )
}
