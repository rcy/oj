import { useAllSpacesQuery } from '../../generated-types';
import { Link } from "react-router-dom";

export default function() {
  const {loading, data } = useAllSpacesQuery()

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
