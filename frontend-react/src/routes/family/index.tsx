import { Routes, Route, Link } from "react-router-dom";
import { useCurrentUserFamilyQuery } from "../../generated-types"
import AdminSection from '../../components/AdminSection';
import Button from '../../Button';

export default function FamilyIndex() {
  const queryResult = useCurrentUserFamilyQuery({ fetchPolicy: 'network-only' })
  const family = queryResult.data?.currentUser?.family;

  if (!family) {
    return null
  }

  return (
    <div className="flex flex-col">

      <header className="flex justify-between px-5 h-20 bg-orange-100 items-center">
        <div className="text-6xl">My Family</div>

        <Link to="/family/add">
          <Button color="blue">add family member</Button>
        </Link>
      </header >

      <main className="flex flex-col mt-5 ml-10 gap-2">
        <Routes>
          <Route path="/" element={
            <div>
              {family.familyMemberships.nodes.map((n) => (
                <Link to={`/family/${n.person?.id}`}>
                  <div className="flex items-center gap-2 hover:bg-red-100">
                    <img src={`${n.person?.avatarUrl}&s=40`} />
                    <div className="text-3xl">
                      {n.person?.name} - {n.role} - {n.person?.user?.id}
                    </div>
                  </div>
                </Link>
              ))}
            </div>
          } />
          <Route path="/add" element={<div>add</div>} />
        </Routes>
      </main>
    </div >
  )
}
