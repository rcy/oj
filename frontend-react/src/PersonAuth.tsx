import { useState } from "react";
import TextInput from "./components/TextInput";
import {
  useCreateLoginCodeMutation,
  useExchangeCodeMutation,
} from "./generated-types";
import { useSearchParams } from "react-router-dom";
import {
  Center,
  Container,
  FormControl,
  FormErrorMessage,
  FormHelperText,
  FormLabel,
  Heading,
  Input,
  VStack,
} from "@chakra-ui/react";

export default function PersonAuth() {
  const [createLoginCode] = useCreateLoginCodeMutation();
  const [exchangeCode] = useExchangeCodeMutation();
  const [searchParams] = useSearchParams();

  const [username, setUsername] = useState("");
  const [loginCodeId, setLoginCodeId] = useState(null);
  const [code, setCode] = useState("");
  const [error, setError] = useState("");

  async function submitUsername(ev: any) {
    ev.preventDefault();
    setError("");
    const result = await createLoginCode({ variables: { username } });
    const newLoginCodeId = result.data?.createLoginCode?.loginCodeId;
    if (!newLoginCodeId) {
      setError(`no such person ${username}`);
      setUsername("");
    } else {
      setLoginCodeId(newLoginCodeId);
    }
  }

  async function submitCode(ev: any) {
    ev.preventDefault();
    setError("");
    const result = await exchangeCode({ variables: { loginCodeId, code } });
    const sessionKey = result.data?.exchangeCode?.sessionKey;
    if (!sessionKey) {
      setError("code does not match");
      setCode("");
    } else {
      localStorage.setItem("sessionKey", sessionKey);
      window.location.assign(searchParams.get("from") || "/");
    }
  }

  if (loginCodeId) {
    return (
      <Center h="100vh">
        <VStack>
          <Container>
            <Heading as="h1" className="text-xl">
              Hello, {username}!
            </Heading>
            <form onSubmit={submitCode}>
              <FormControl isInvalid={error.length > 0}>
                <FormLabel>
                  Enter the code sent to your parent's email
                </FormLabel>
                <Input
                  name="code"
                  onChange={(ev: React.ChangeEvent<HTMLInputElement>) => {
                    setCode(ev.target.value);
                  }}
                  value={code}
                />
                <FormErrorMessage>{error}</FormErrorMessage>
              </FormControl>
            </form>
          </Container>
        </VStack>
      </Center>
    );
  }

  return (
    <Center h="100vh">
      <form onSubmit={submitUsername}>
        <VStack>
          <Container>
            <Heading as="h1" className="text-xl">
              Welcome!
            </Heading>
            <FormControl isInvalid={error.length > 0}>
              <FormLabel>What is your username?</FormLabel>
              <Input
                name="username"
                onChange={(ev: React.ChangeEvent<HTMLInputElement>) =>
                  setUsername(ev.target.value)
                }
                value={username}
                autoComplete="off"
              />
              <FormErrorMessage>{error}</FormErrorMessage>
            </FormControl>
          </Container>
        </VStack>
      </form>
    </Center>
  );
}
