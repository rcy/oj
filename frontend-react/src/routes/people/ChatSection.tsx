import { useContext, useState } from "react";
import Chat from "../../components/Chat";
import { PersonIdContext } from "../../contexts";
import {
  useCreateSpaceMembershipMutation,
  useCreateSpaceMutation,
  useSharedSpacesQuery,
} from "../../generated-types";

export default function ChatSection({ pagePerson }: any) {
  const myPersonId = useContext(PersonIdContext);
  const [createSpace] = useCreateSpaceMutation();
  const [createSpaceMembership] = useCreateSpaceMembershipMutation();
  const sharedSpacesQuery = useSharedSpacesQuery({
    variables: { person1: pagePerson?.id, person2: myPersonId },
  });
  const [creating, setCreating] = useState(false);

  async function handleClickCreateSpace() {
    setCreating(true);

    // create the space
    let result = await createSpace({
      variables: { name: "dm" },
    });
    if (result.errors) {
      throw result.errors;
    }
    const spaceId = result.data?.createSpace?.space?.id;

    // add the page person
    result = await createSpaceMembership({
      variables: { spaceId, personId: pagePerson?.id },
    });
    if (result.errors) {
      throw result.errors;
    }

    // add self
    result = await createSpaceMembership({
      variables: { spaceId, personId: myPersonId },
    });
    if (result.errors) {
      throw result.errors;
    }

    await sharedSpacesQuery.refetch();
    setCreating(false);
  }

  if (sharedSpacesQuery.loading) {
    return <div>loading</div>;
  }

  if (creating) {
    return <div>creating</div>;
  }

  const spaces = sharedSpacesQuery.data?.spaces?.edges || [];

  if (spaces.length === 0) {
    return (
      <div>
        <button onClick={handleClickCreateSpace}>create space to chat</button>
      </div>
    );
  }

  return (
    <div>
      <Chat spaceId={spaces[0].node.id} />
    </div>
  );
}
