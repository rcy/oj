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
    <div className="grid h-screen place-items-center bg-orange-200">
      <div>
        <h1 className="pb-20 text-xl">Select Family Member</h1>
        <section className="flex space-x-10">
          {data?.currentUser?.family?.familyMemberships.nodes.map((m: any) => (
            <a
              href="#"
              onClick={(ev) => become(ev, m.id)}
              key={m.id}
              className="rounded-lg bg-blue-200 w-48 text-center"
            >
              <img className="rounded-t-lg" src={m.person.avatarUrl} alt=""/>

              <div className="p-2">
                {m.person.name}
              </div>
            </a>
          ))}
        </section>
      </div>
    </div>
  )
}
