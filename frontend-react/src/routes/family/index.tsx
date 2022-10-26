import { useCurrentUserFamilyQuery } from "../../generated-types"

export default function FamilyIndex() {
  const queryResult = useCurrentUserFamilyQuery()
  const family = queryResult.data?.currentUser?.family;

  if (!family) {
    return null
  }

  return (
    <div className="p-10">
      <div className="flex gap-2">
        {family.familyMemberships.nodes.map((n) => (
          <div>
            <img src={n.person?.avatarUrl} />
            {n.person?.name}
          </div>
        ))}
      </div>
    </div>
  )
}
