import { MouseEventHandler } from 'react';
import { useCurrentFamilyMembershipQuery } from './generated-types';
import { useNavigate } from 'react-router-dom';
import { Routes, Route } from "react-router-dom";
import MemberHome from './MemberHome';
import AdminLayout from './admin/AdminLayout';
import PageNotFound from './PageNotFound';

import SpacesIndex from './routes/spaces/index';
import SpacesShow from './routes/spaces/show';
import { PersonIdContext } from './contexts'

type MemberLayoutType = { doLogout: Function }

export default function MemberLayout({ doLogout }: MemberLayoutType) {
  const { loading, data } = useCurrentFamilyMembershipQuery({ fetchPolicy: 'network-only' })
  let navigate = useNavigate();

  if (loading) { return <span>loading</span> }

  const handleLogout: MouseEventHandler = (ev) => {
    ev.preventDefault()
    navigate('/');
    doLogout()
  }

  return (
    <PersonIdContext.Provider value={data?.currentFamilyMembership?.person?.id}>
      <div className="font-sans">
        <nav className="bg-gray-800 px-2 py-2 text-white flex justify-between text-xl">
          <div className="flex items-center space-x-2">
            <div>🐙</div><a href="#" onClick={handleLogout}>Octopus Jr</a>
          </div>
          <div className="text-orange-200">
            {data?.currentFamilyMembership?.person?.name} ({data?.currentFamilyMembership?.role})
          </div>
        </nav>
        <main>

          <Routes>
            <Route path="/" element={<MemberHome familyMembership={data?.currentFamilyMembership} />} />
            <Route path="/spaces" element={<SpacesIndex />} />
            <Route path="/spaces/:id" element={<SpacesShow />} />
            <Route path="/admin/*" element={<AdminLayout />} />
            <Route path="*" element={<PageNotFound />} />
          </Routes>
        </main>
      </div>
    </PersonIdContext.Provider>
  )
}
