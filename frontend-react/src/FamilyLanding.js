import React from 'react'
import { CURRENT_USER_FAMILY } from './queries.js'
import { useQuery } from '@apollo/client';
import { useNavigate } from 'react-router-dom';

export default function FamilyLanding() {
  const { loading, error, data } = useQuery(CURRENT_USER_FAMILY)
  let navigate = useNavigate();

  if (loading) { return "loading" }

  function addFamilyMember() {
    alert('add family member')
  }

  function become(membershipId) {
    console.log({ membershipId })
    window.sessionStorage.setItem('membershipId', membershipId)
    navigate('/member')
  }

  return (
    <section>
      <h1>family</h1>
      <h2>members</h2>
      {data.currentUser.family.familyMemberships.nodes.map((m) => (
        <div key={m.id}>
          {m.person.name} ({m.role}) <button onClick={() => become(m.id)}>become</button>
        </div>
      ))}
    </section>
  )
}
