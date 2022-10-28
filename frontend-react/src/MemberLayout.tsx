import { useCurrentFamilyMembershipQuery } from './generated-types';
import { Link, Navigate, NavLink } from 'react-router-dom';
import { Routes, Route } from "react-router-dom";
import { PersonIdContext } from './contexts'
import { ReactNode } from 'react';
import MemberHome from './MemberHome';
import AdminLayout from './admin/AdminLayout';
import PageNotFound from './PageNotFound';

import SpacesIndex from './routes/spaces/index';
import SpacesShow from './routes/spaces/show';
import SpacesExplore from './routes/spaces/explore';
import FamilyIndex from './routes/family/index';
import HackIndex from './routes/hack/index';

type MemberLayoutType = { doLogout: Function }

function NavCell({children}:any) {
  return (
    <div className='mb-5'>
      {children}
    </div>
  )
}

export default function MemberLayout({ doLogout }: MemberLayoutType) {
  const { loading, data } = useCurrentFamilyMembershipQuery({ fetchPolicy: 'network-only' })
  if (loading) { return <span>loading</span> }

  return (
    <PersonIdContext.Provider value={data?.currentFamilyMembership?.person?.id}>
      <div className="font-sans">
        <nav className="text-black flex gap-2">

          <aside className="flex-none flex flex-col items-center w-21 text-white">
            <NavCell>
              <MyNavLink to="/me" activeClass="bg-red-200 text-black">
                <img width="80" src={data?.currentFamilyMembership?.person?.avatarUrl} />
              </MyNavLink>
            </NavCell>
            <NavCell>
              <MyNavLink to="/family" activeClass="bg-blue-200 text-black">
                <img width="80" src="https://www.gravatar.com/avatar/07ae617e?f=y&d=identicon" />
              </MyNavLink>
            </NavCell>
            <NavCell>
              <MyNavLink to="/spaces" activeClass="bg-green-200 text-black">
                <div className="text-6xl pb-2">🧭</div>
              </MyNavLink>
            </NavCell>
            <NavCell>
              <Link to="/xyz">
                <img height="80" width="80" src="https://tse4.mm.bing.net/th?id=OIP.fbwROgl5jUOKwUD2XkCXRAHaGw&pid=Api" />
              </Link>
            </NavCell>
          </aside>

          <main className='w-full'>
            <Routes>
              <Route path="/" element={<Navigate to="/me" />} />
              <Route path="/me" element={<MemberHome familyMembership={data?.currentFamilyMembership} doLogout={doLogout} />} />
              <Route path="/spaces" element={<SpacesIndex />} />
              <Route path="/family" element={<FamilyIndex />} />
              <Route path="/spaces/explore" element={<SpacesExplore />} />
              <Route path="/spaces/:id" element={<SpacesShow />} />
              <Route path="/admin/*" element={<AdminLayout />} />
              <Route path="/hack/*" element={<HackIndex />} />
              <Route path="*" element={<PageNotFound />} />
            </Routes>
          </main>
        </nav>
      </div>
    </PersonIdContext.Provider >
  )
}

interface MyNavLinkProps {
  to: string,
  activeClass: string,
  children: ReactNode,
}
function MyNavLink({ to, children, activeClass }: MyNavLinkProps) {
  return (
    <NavLink
      to={to}
      className={({ isActive }) => isActive ? activeClass : ""}
    >
      {children}
    </NavLink>
  )
}
