import { useQuery } from '@apollo/client';
import { graphql } from '../../gql';
import { Link } from "react-router-dom";

const ALL_SPACES = graphql(`
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
`);

export default function() {
  const {loading, data } = useQuery(ALL_SPACES)

  if (loading) { return null }

  return (
    <div className="flex flex-col">
      {data?.spaces?.edges.map((s) => (
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
