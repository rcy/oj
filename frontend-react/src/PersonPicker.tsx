import { MouseEvent } from "react";
import { useCurrentUserWithManagedPeopleQuery } from "./generated-types";

type PersonPickerType = { setPersonId: Function };

export default function PersonPicker({ setPersonId }: PersonPickerType) {
  const { loading, data } = useCurrentUserWithManagedPeopleQuery();

  if (loading) {
    return null;
  }

  const managedPeople = data?.currentUser?.managedPeople;

  function become(ev: MouseEvent, id: string) {
    ev.preventDefault();
    setPersonId(id);
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
      <main className="py-16 w-full flex flex-col items-center gap-2">
        <PersonCard person={data?.currentUser?.person} onClick={become} />

        <div>managed accounts:</div>
        {managedPeople?.nodes.map((m: any) => (
          <PersonCard key={m.id} person={m.person} onClick={become} />
        ))}
      </main>
    </div>
  );
}

interface PersonCardProps {
  person: any;
  onClick: Function;
}
function PersonCard({ person, onClick }: PersonCardProps) {
  return (
    <a
      href={`#${person.id}`}
      onClick={(ev) => onClick(ev, person.id)}
      key={person.id}
      className="rounded-lg w-72 flex bg-white items-center border-solid border-4 hover:border-orange-600"
    >
      <img className="rounded-l-sm" src={person.avatarUrl} alt={person.name} />

      <div className="p-3 text-xl">{person.name}</div>
    </a>
  );
}
