import { useState } from 'react';
import { useMutation } from '@apollo/client';
import { CREATE_SPACE } from '../queries.js';
import TextInput from './TextInput.js';
import Button from '../Button';

export default function AdminAddFamilyMember() {
  const [value, setValue] = useState('')
  const [createSpace, { data, loading, error }] = useMutation(CREATE_SPACE);

  async function handleSubmit(ev) {
    ev.preventDefault()
    console.log('mutating...')
    const result = await createSpace({
      variables: { name: value }
    })
    console.log('mutating...done', { result })
  }

  return (
    <>
      <h1 className="text-xl pb-10">Create Space</h1>

      <form onSubmit={handleSubmit}>
        <TextInput
          label="Space Name"
          name="name"
          onChange={(ev) => setValue(ev.target.value)}
          value={value}
        />

        <Button type='submit'>Submit</Button>
      </form>
    </>
  )
}
