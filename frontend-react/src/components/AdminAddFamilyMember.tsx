import { FormEventHandler, useState } from "react";
import { useCreateNewFamilyMemberMutation } from "../generated-types";
import TextInput from "./TextInput";
import { Box, Button, ButtonGroup, Flex, FormControl, FormHelperText, FormLabel, Heading, Input, Text, VStack } from "@chakra-ui/react";

interface Props {
  onSuccess: Function;
  onCancel: Function;
}

export default function AdminAddFamilyMember({ onSuccess, onCancel }: Props) {
  const [firstName, setFirstName] = useState("");
  const [addFamilyMember] = useCreateNewFamilyMemberMutation();

  const handleSubmit: FormEventHandler = async (ev) => {
    ev.preventDefault();
    const trimmedValue = firstName.trim();
    if (trimmedValue.length) {
      console.log("mutating...");
      const result = await addFamilyMember({
        variables: { name: trimmedValue, role: "child" },
      });
      console.log("mutating...done", { result });
      if (!result.errors) {
        onSuccess(
          result.data?.createNewFamilyMember?.familyMembership?.personId
        );
      }
    }
  };

  const handleCancel = () => {
    onCancel();
  };

  return (
    <VStack align="stretch">
      <Heading size="lg" mb="4">Create A Managed Account For Your Child</Heading>
      <form onSubmit={handleSubmit}>
        <VStack align="stretch" spacing={10}>
          <FormControl>
            <FormLabel>What is your child's name?</FormLabel>
            <Input
              name="name"
              onChange={(ev: React.ChangeEvent<HTMLInputElement>) =>
                setFirstName(ev.target.value)
              }
              value={firstName}
            />
            <FormHelperText>First name only is fine</FormHelperText>
          </FormControl>

          <ButtonGroup>
            <Button colorScheme="green" type="submit">
              Submit
            </Button>

            <Button onClick={handleCancel}>
              Cancel
            </Button>
          </ButtonGroup>
        </VStack>
      </form>
    </VStack>
  );
}
