import { MouseEvent } from 'react'
import { useCurrentUserFamilyQuery } from './generated-types';
import { useQuery } from '@apollo/client';

type FamilyLandingType = { setFamilyMembershipId: Function }

export default function FamilyLanding({ setFamilyMembershipId }: FamilyLandingType) {
  const { loading, data } = useCurrentUserFamilyQuery();

  if (loading) { return null }

  function become(ev: MouseEvent, id: string) {
    ev.preventDefault()
    setFamilyMembershipId(id)
  }

  return (
    <div className="grid bg-gradient-to-b from-orange-300 to-yellow-300">
      <section className="py-10 w-full flex flex-col items-center gap-5">
        {data?.currentUser?.family?.familyMemberships.nodes.map((m: any) => (
          <a
            href="#"
            onClick={(ev) => become(ev, m.id)}
            key={m.id}
            className="rounded-lg w-64 flex bg-white items-center border-solid border-4 hover:border-orange-600"
          >
            <img className="rounded-l-sm" src={m.person.avatarUrl} alt={m.person.name} />

            <div className="p-3 text-xl">
              {m.person.name}
            </div>
          </a>
        ))}
      </section>
    </div>
  )
}
