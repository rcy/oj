import { useCurrentPersonQuery } from "../generated-types";
import { Navigate, NavLink } from "react-router-dom";
import { Routes, Route } from "react-router-dom";
import { PersonIdContext } from "../contexts";
import { ReactNode } from "react";
import MeIndex from "./me/index";
import AdminLayout from "../components/AdminLayout";
import PageNotFound from "../PageNotFound";

import SpacesIndex from "./spaces/index";
import SpacesShow from "./spaces/show";
import SpacesExplore from "./spaces/explore";
import FamilyIndex from "./family/index";
import MemberSetProfilePicture from "./member/MemberSetProfilePicture";
import PeopleIndexPage from "./people";

type MemberLayoutType = { doLogout: Function };

function NavCell({ children }: any) {
  return <div className="mb-5">{children}</div>;
}

export default function MemberLayout({ doLogout }: MemberLayoutType) {
  const { loading, data } = useCurrentPersonQuery({
    fetchPolicy: "network-only",
  });
  if (loading) {
    return <span>loading</span>;
  }

  return (
    <PersonIdContext.Provider value={data?.currentPerson?.id}>
      <div className="font-sans">
        <nav className="text-black flex gap-2">
          <aside className="flex-none flex flex-col items-center w-21 text-white p-2 border-solid border-r-2 border-gray-300">
            <NavCell>
              <MyNavLink to="/me" inactiveClassName="" activeClass="">
                <img
                  alt="avatar"
                  width="80"
                  src={data?.currentPerson?.avatarUrl}
                />
              </MyNavLink>
            </NavCell>
          </aside>

          <main className="w-full">
            <Routes>
              <Route path="/" element={<Navigate to="/me" />} />
              <Route
                path="/me"
                element={
                  <MeIndex
                    doLogout={doLogout}
                  />
                }
              />
              <Route path="/me/pic" element={<MemberSetProfilePicture />} />
              <Route path="/spaces" element={<SpacesIndex />} />
              <Route path="/family/*" element={<FamilyIndex />} />
              <Route path="/people/*" element={<PeopleIndexPage />} />
              <Route path="/spaces/explore" element={<SpacesExplore />} />
              <Route path="/spaces/:id" element={<SpacesShow />} />
              <Route path="/admin/*" element={<AdminLayout />} />
              <Route path="*" element={<PageNotFound />} />
            </Routes>
          </main>
        </nav>
      </div>
    </PersonIdContext.Provider>
  );
}

interface MyNavLinkProps {
  to: string;
  inactiveClassName?: string;
  activeClass: string;
  children: ReactNode;
}
function MyNavLink({
  to,
  children,
  inactiveClassName,
  activeClass,
}: MyNavLinkProps) {
  return (
    <NavLink
      to={to}
      className={({ isActive }) => (isActive ? activeClass : inactiveClassName)}
    >
      {children}
    </NavLink>
  );
}
