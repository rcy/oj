import { Routes, Route, Link, useNavigate } from "react-router-dom";
import {
  useCurrentFamilyMembershipQuery,
  useCurrentUserFamilyQuery,
} from "../../generated-types";
import Button from "../../Button";
import AdminAddFamilyMember from "../../components/AdminAddFamilyMember";
import MemberPageNotFound from "../../PageNotFound";
import { useContext } from "react";
import { PersonIdContext } from "../../contexts";

export default function FamilyIndex() {
  const queryResult = useCurrentUserFamilyQuery({
    fetchPolicy: "network-only",
  });
  const currentUserFamilyMembershipQuery = useCurrentFamilyMembershipQuery();
  const family = queryResult.data?.currentUser?.family;
  const navigate = useNavigate();
  const personId = useContext(PersonIdContext);

  if (!family) {
    return null;
  }

  const role =
    currentUserFamilyMembershipQuery.data?.currentFamilyMembership?.role;

  async function handleAddSuccess(personId: string) {
    await queryResult.refetch();
    navigate("/people/" + personId);
  }

  return (
    <div className="flex flex-col">
      <header className="flex justify-between items-center">
        <div className="text-4xl">Family</div>

        {role === "admin" && (
          <Link to="/family/add">
            <Button color="blue">add family member</Button>
          </Link>
        )}
      </header>

      <main className="flex flex-col mt-5 gap-2">
        <Routes>
          <Route
            path="/"
            element={
              <div>
                {family.familyMemberships.nodes.map((n) => (
                  <div key={n.person?.id}>
                    <Link
                      to={
                        n.person?.id === personId
                          ? "/me"
                          : `/people/${n.person?.id}`
                      }
                    >
                      <div className="flex items-center gap-2">
                        <img alt="avatar" src={`${n.person?.avatarUrl}&s=40`} />
                        <div className="text-3xl">
                          {n.title || n.person?.name}
                        </div>
                      </div>
                    </Link>
                  </div>
                ))}
              </div>
            }
          />
          <Route
            path="/add"
            element={
              <AdminAddFamilyMember
                onSuccess={handleAddSuccess}
                onCancel={() => navigate("/me")}
              />
            }
          />
          <Route path="*" element={<MemberPageNotFound />} />
        </Routes>
      </main>
    </div>
  );
}
