import { Link } from "react-router-dom";
import { useCurrentUserFamilyQuery } from "../../generated-types"

export default function FamilyIndex() {
  const queryResult = useCurrentUserFamilyQuery()
  const family = queryResult.data?.currentUser?.family;

  if (!family) {
    return null
  }

  return (
    <div className="flex flex-col">
      <div className="h-20 text-6xl bg-orange-100">My Family</div>

      <div className="flex flex-col mt-5 ml-10 gap-2">
        {family.familyMemberships.nodes.map((n) => (
          <Link to={`/family/${n.person?.id}`}>
            <div className="flex items-center gap-2 hover:bg-red-100">
              <img src={`${n.person?.avatarUrl}&s=40`} />
              <div className="text-3xl">
                {n.person?.name}
              </div>
            </div>
          </Link>
        ))}
      </div>
    </div>
  )
}
