import { MouseEvent } from "react";
import { useCurrentUserFamilyQuery } from "./generated-types";

type FamilyLandingType = { setFamilyMembershipId: Function };

export default function FamilyLanding({
  setFamilyMembershipId,
}: FamilyLandingType) {
  const { loading, data } = useCurrentUserFamilyQuery();

  if (loading) {
    return null;
  }

  const familyMemberships = data?.currentUser?.family?.familyMemberships;

  function become(ev: MouseEvent, id: string) {
    ev.preventDefault();
    setFamilyMembershipId(id);
  }

  return (
    <div className="bg-gradient-to-b from-orange-300 to-yellow-300 min-h-screen flex flex-col">
      <header className="pr-2 flex justify-between bg-orange-200">
        <div className="flex gap-1">
          <img width="32" src="octopus1.png" alt="Octopus" />
          Octopus Jr.
        </div>
        <div>{data?.currentUser?.name}'s Account</div>
      </header>
      <main className="py-16 w-full flex flex-col items-center gap-5">
        {familyMemberships?.nodes.map((m: any) => (
          <a
            href={`#${m.id}`}
            onClick={(ev) => become(ev, m.id)}
            key={m.id}
            className="rounded-lg w-72 flex bg-white items-center border-solid border-4 hover:border-orange-600"
          >
            <img
              className="rounded-l-sm"
              src={m.person.avatarUrl}
              alt={m.person.name}
            />

            <div className="p-3 text-xl">{m.person.name}</div>
          </a>
        ))}
      </main>
    </div>
  );
}
