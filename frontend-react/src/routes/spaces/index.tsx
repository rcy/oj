import { useQuery, gql } from '@apollo/client';
import { Link } from "react-router-dom";

const ALL_SPACES = gql`
  query AllSpaces {
    spaces {
      edges {
        node {
          id
          name
        }
      }
    }
  }
`

export default function() {
  const { loading, data } = useQuery(ALL_SPACES)

  console.log({ data })

  if (loading) { return null }

  return (
    <div className="flex flex-col">
      {data.spaces.edges.map((s: any) => (
        <Link
          to={`/spaces/${s.node.id}`}
          key={s.node.id}
        >
          {s.node.name}
        </Link>
      ))}
    </div>
  );
}
