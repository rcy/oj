import React from 'react'
import { CURRENT_FAMILY_MEMBERSHIP } from './queries.js'
import { useQuery } from '@apollo/client';
import { useNavigate } from 'react-router-dom';
import AdminSection from './AdminSection.js';
import { clearFamilyMembershipId } from './util/family.js';

export default function MemberHome({ logout }) {
  const { loading, error, data } = useQuery(CURRENT_FAMILY_MEMBERSHIP, { fetchPolicy: 'network-only' })
  let navigate = useNavigate();

  if (loading) { return "loading" }

  console.log({ loading, error, data })

  function addFamilyMember() {
    // https://www.apollographql.com/docs/react/data/mutations
    alert('not implemented')
  }

  return (
    <div className="font-sans">
      <nav className="bg-gray-800 px-2 py-2 text-white flex justify-between text-xl">
        <div className="flex items-center space-x-2">
          <div>🐱</div><a href="#" onClick={logout}>Small Space</a>
        </div>
        <div className="text-orange-200">
          {data.currentFamilyMembership.person.name} ({data.currentFamilyMembership.role})
        </div>
      </nav>
      <main className="p-10">
        {data.currentFamilyMembership.role === 'admin' && <AdminSection />}
      </main>
    </div>
  )
}
