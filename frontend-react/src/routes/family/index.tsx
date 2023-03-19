import { Routes, Route, Link, useNavigate } from "react-router-dom";
import { useCurrentPersonFamilyMembershipQuery } from "../../generated-types";
import Button from "../../Button";
import AdminAddFamilyMember from "../../components/AdminAddFamilyMember";
import MemberPageNotFound from "../../PageNotFound";
import { useContext } from "react";
import { PersonIdContext } from "../../contexts";
import { graphql } from "../../gql";

graphql(`
  fragment FamilyMembershipItem on FamilyMembership {
    id
    role
    title
    person {
      id
      name
      avatarUrl
      username
      user {
        id
      }
    }
  }
`);

graphql(`
  query CurrentPersonFamilyMembership {
    currentPerson {
      id
      familyMembership {
        id
        role
        family {
          id
          familyMemberships(orderBy: CREATED_AT_ASC) {
            edges {
              node {
                ...FamilyMembershipItem
              }
            }
          }
        }
      }
    }
  }
`);

export default function FamilyIndex() {
  const query = useCurrentPersonFamilyMembershipQuery();
  const family = query.data?.currentPerson?.familyMembership?.family;
  const navigate = useNavigate();
  const personId = useContext(PersonIdContext);

  if (!family) {
    return null;
  }

  const role = query.data?.currentPerson?.familyMembership?.role;

  async function handleAddSuccess(personId: string) {
    await query.refetch();
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
                {family.familyMemberships.edges.map(
                  ({ node: familyMembership }) => (
                    <div key={familyMembership.id}>
                      <Link
                        to={
                          familyMembership.person?.id === personId
                            ? "/me"
                            : `/people/${familyMembership.person?.id}`
                        }
                      >
                        <div className="flex items-center gap-2">
                          <img
                            alt="avatar"
                            src={`${familyMembership.person?.avatarUrl}&s=40`}
                          />
                          <div className="text-3xl">
                            {familyMembership.title ||
                              familyMembership.person?.name}{" "}
                            {familyMembership.person?.user?.id
                              ? null
                              : "managed account"}
                          </div>
                        </div>
                      </Link>
                    </div>
                  )
                )}
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
