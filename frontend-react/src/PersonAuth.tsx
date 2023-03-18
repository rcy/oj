import { useState } from "react";
import TextInput from "./components/TextInput";
import { useCreateLoginCodeMutation, useExchangeCodeMutation } from "./generated-types";
import { useSearchParams } from "react-router-dom";

export default function PersonAuth() {
  const [createLoginCode] = useCreateLoginCodeMutation()
  const [exchangeCode] = useExchangeCodeMutation()
  const [searchParams] = useSearchParams()

  const [username, setUsername] = useState("")
  const [loginCodeId, setLoginCodeId] = useState(null)
  const [code, setCode] = useState("")
  const [error, setError] = useState("")

  async function submitUsername(ev: any) {
    ev.preventDefault()
    setError('')
    const result = await createLoginCode({ variables: { username } })
    const newLoginCodeId = result.data?.createLoginCode?.loginCodeId;
    if (!newLoginCodeId) {
      setError(`no such person ${username}`)
      setUsername('')
    } else {
      setLoginCodeId(newLoginCodeId)
    }
  }

  async function submitCode(ev: any) {
    ev.preventDefault()
    setError('')
    const result = await exchangeCode({ variables: { loginCodeId, code } })
    const sessionKey = result.data?.exchangeCode?.sessionKey
    if (!sessionKey) {
      setError('code does not match')
      setCode('')
    } else {
      localStorage.setItem('sessionKey', sessionKey)
      window.location.assign(searchParams.get('from') || '/')
    }
  }

  if (loginCodeId) {
    return (
      <div className="grid h-screen place-items-center">
        <div className="flex flex-col items-center">
          <h1 className="text-xl">hello {username}</h1>
          <form onSubmit={submitCode}>
            <TextInput
              label="Enter the code"
              name="code"
              onChange={(ev: React.ChangeEvent<HTMLInputElement>) => { setCode(ev.target.value) }}
              value={code}
            />
            {error}
          </form>
        </div>
      </div>
    )
  }

  return (
    <div className="grid h-screen place-items-center">
      <div className="flex flex-col items-center">
        <form onSubmit={submitUsername}>
          <TextInput
            label="Enter your username"
            name="username"
            onChange={(ev: React.ChangeEvent<HTMLInputElement>) => setUsername(ev.target.value)}
            value={username}
          />
          {error}
        </form>
      </div>
    </div>
  )
}
