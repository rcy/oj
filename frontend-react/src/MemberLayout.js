import React from 'react'
import { CURRENT_FAMILY_MEMBERSHIP } from './queries.js'
import { useQuery } from '@apollo/client';
import { useNavigate } from 'react-router-dom';
import { clearFamilyMembershipId } from './util/family.js';
import { Routes, Route, Link, navigate } from "react-router-dom";
import MemberHome from './MemberHome.js';
import Button from './Button.js';
import AdminLayout from './admin/AdminLayout.js';
import PageNotFound from './PageNotFound.js';

export default function MemberLayout({ doLogout }) {
  const { loading, error, data } = useQuery(CURRENT_FAMILY_MEMBERSHIP, { fetchPolicy: 'network-only' })
  let navigate = useNavigate();

  if (loading) { return "loading" }

  function handleLogout(ev) {
    ev.preventDefault()
    navigate('/');
    doLogout()
  }

  function addFamilyMember() {
    // https://www.apollographql.com/docs/react/data/mutations
    alert('not implemented')
  }

  return (
    <div className="font-sans">
      <nav className="bg-gray-800 px-2 py-2 text-white flex justify-between text-xl">
        <div className="flex items-center space-x-2">
          <div>🐙</div><a href="#" onClick={handleLogout}>Octopus Jr</a>
        </div>
        <div className="text-orange-200">
          {data.currentFamilyMembership.person.name} ({data.currentFamilyMembership.role})
        </div>
      </nav>
      <main className="p-10">
        <Routes>
          <Route path="/" element={<MemberHome familyMembership={data.currentFamilyMembership} />}/>
          <Route path="/admin/*" element={<AdminLayout />}/>
          <Route path="*" element={<PageNotFound />} />
        </Routes>
      </main>
    </div>
  )
}
