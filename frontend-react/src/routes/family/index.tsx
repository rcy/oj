import { Routes, Route, Link, useNavigate } from "react-router-dom";
import { useCurrentUserFamilyQuery } from "../../generated-types"
import Button from '../../Button';
import AdminAddFamilyMember from "../../components/AdminAddFamilyMember";
import MemberPageNotFound from "../../PageNotFound";

export default function FamilyIndex() {
  const queryResult = useCurrentUserFamilyQuery({ fetchPolicy: 'network-only' })
  const family = queryResult.data?.currentUser?.family;
  const navigate = useNavigate();

  if (!family) {
    return null
  }

  async function handleAddSuccess(personId: string) {
    await queryResult.refetch()
    navigate('/people/'+personId)
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
                <Link key={n.person?.id} to={`/people/${n.person?.id}`}>
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
          <Route
            path="/add"
            element={<AdminAddFamilyMember onSuccess={handleAddSuccess} onCancel={() => navigate('/family')} />}
          />
          <Route path="*" element={<MemberPageNotFound/>}/>
        </Routes>
      </main>
    </div >
  )
}
