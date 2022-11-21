import { useContext } from "react";
import Button from "../Button";
//import { useJoinSpaceMutation } from '../generated-types'
import { PersonIdContext } from "../contexts";
import { useJoinSpaceMutation } from "../generated-types";

interface Props {
  spaceId: string;
}

export default function JoinSpaceButton({ spaceId }: Props) {
  const personId = useContext(PersonIdContext);
  const [joinSpace] = useJoinSpaceMutation({
    variables: { spaceId, personId },
  });

  const handleClickJoin = async () => {
    console.log("join", spaceId, personId);
    const result = await joinSpace();
    console.log({ result });
  };

  return <Button onClick={handleClickJoin}>join space</Button>;
}
