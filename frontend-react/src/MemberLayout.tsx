import { MouseEventHandler } from 'react'
import { CURRENT_FAMILY_MEMBERSHIP } from './queries'
import { useQuery } from '@apollo/client';
import { useNavigate } from 'react-router-dom';
import { Routes, Route } from "react-router-dom";
import MemberHome from './MemberHome';
import AdminLayout from './admin/AdminLayout';
import PageNotFound from './PageNotFound';

import SpacesIndex from './routes/spaces/index';
import SpacesShow from './routes/spaces/show';

type MemberLayoutType = { doLogout: Function }

export default function MemberLayout({ doLogout }: MemberLayoutType) {
  const { loading, data } = useQuery(CURRENT_FAMILY_MEMBERSHIP, { fetchPolicy: 'network-only' })
  let navigate = useNavigate();

  if (loading) { return <span>loading</span> }

  const handleLogout: MouseEventHandler = (ev) => {
    ev.preventDefault()
    navigate('/');
    doLogout()
  }

  return (
    <div className="font-sans">
      <nav className="bg-gray-800 px-2 py-2 text-white flex justify-between text-xl">
        <div className="flex items-center space-x-2">
          <div>🐙</div><a href="#" onClick={handleLogout}>Octopus Jr</a>
        </div>
        <div className="text-orange-200">
          {data?.currentFamilyMembership?.person?.name} ({data?.currentFamilyMembership?.role})
        </div>
      </nav>
      <main className="p-10">
        <Routes>
          <Route path="/" element={<MemberHome familyMembership={data?.currentFamilyMembership} />}/>
          <Route path="/spaces" element={<SpacesIndex />}/>
          <Route path="/spaces/:id" element={<SpacesShow />}/>
          <Route path="/admin/*" element={<AdminLayout />}/>
          <Route path="*" element={<PageNotFound />} />
        </Routes>
      </main>
    </div>
  )
}
