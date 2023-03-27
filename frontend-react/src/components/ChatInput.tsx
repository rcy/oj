import { Avatar, Input, InputGroup, InputLeftElement } from "@chakra-ui/react";
import { ChangeEventHandler, KeyboardEventHandler, useState } from "react";
import { Person } from "../generated-types";

interface Props {
  onSubmit: Function;
  person: any
}

export default function ChatInput({ onSubmit, person }: Props) {
  const [input, setInput] = useState("");

  const handleChange: ChangeEventHandler<HTMLInputElement> = (ev) => {
    setInput(ev.target.value);
  };

  const handleKey: KeyboardEventHandler<HTMLInputElement> = (ev) => {
    if (ev.key === "Enter") {
      if (input.length > 0 && onSubmit(input)) {
        setInput("");
      }
    }
  };

  return (
    <InputGroup>
      <InputLeftElement>
        <Avatar size="xs" name={person.name} src={person.avatarUrl} />
      </InputLeftElement>
      <Input
        background="white"
        type="text"
        placeholder="say something"
        onChange={handleChange}
        onKeyDown={handleKey}
        value={input}
      />
    </InputGroup>
  );
}
