import { useContext } from 'react';
import { PersonIdContext } from '../../contexts';
import { useSpaceMembershipsByPersonIdQuery } from "../../generated-types";
import Debug from "../../components/Debug";

export default function () {
  const personId = useContext(PersonIdContext)

  const queryResult = useSpaceMembershipsByPersonIdQuery({
    variables: {
      personId
    }
  })

  return (
    <div>
      {queryResult.data?.spaceMemberships?.edges.map(({node: m}) => (
        <div key={m.id}>
          {m.space?.name}
        </div>
      ))}
    </div>
  )
}
