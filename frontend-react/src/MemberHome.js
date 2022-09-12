import React from 'react'
import { CURRENT_FAMILY_MEMBERSHIP } from './queries.js'
import { useQuery } from '@apollo/client';
import { useNavigate } from 'react-router-dom';

export default function MemberHome() {
  const { loading, error, data } = useQuery(CURRENT_FAMILY_MEMBERSHIP, { fetchPolicy: 'network-only' })
  let navigate = useNavigate();

  if (loading) { return "loading" }

  function addFamilyMember() {
    // https://www.apollographql.com/docs/react/data/mutations
    alert('not implemented')
  }

  return (
    <section>
      <h1>{data.currentFamilyMembership.person.name} ({data.currentFamilyMembership.role})</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre>

      {data.currentFamilyMembership.role === 'admin' && <button onClick={addFamilyMember}>add family member</button>}
    </section>
  )
}
