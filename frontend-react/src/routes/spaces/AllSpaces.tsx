import { useAllSpacesQuery } from '../../generated-types';
import { Link } from "react-router-dom";
import JoinSpaceButton from '../../components/JoinSpaceButton';
import { useContext } from 'react';
import { PersonIdContext } from '../../contexts';

export default function AllSpaces() {
  const personId = useContext(PersonIdContext)

  const allSpacesResult = useAllSpacesQuery({
    variables: {
      personId
    }
  })

  return (
    <div>
      {allSpacesResult.data?.spaces?.edges.map(({ node: space }) => (
        <div className="flex justify-between" key={space.id}>
          <Link to={`/spaces/${space.id}`}>
            {space.name}
          </Link>
          {space.spaceMemberships.edges.find(x => x)?.node.id ? 'joined' : <JoinSpaceButton spaceId={space.id} />}
        </div>
      ))
      }
    </div>
  )
}
