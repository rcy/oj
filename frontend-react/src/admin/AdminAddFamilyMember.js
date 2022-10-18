import { useState } from 'react';
import { useMutation } from '@apollo/client';
import { CREATE_NEW_FAMILY_MEMBER } from '../queries.js';
import TextInput from './TextInput.js';
import Button from '../Button';

export default function AdminAddFamilyMember() {
  const [value, setValue] = useState('')
  const [addFamilyMember, { data, loading, error }] = useMutation(CREATE_NEW_FAMILY_MEMBER);

  async function handleSubmit(ev) {
    ev.preventDefault()
    console.log('mutating...')
    const result = await addFamilyMember({
      variables: { name: value, role: 'child' }
    })
    console.log('mutating...done', { result })
  }

  return (
    <>
      <h1 className="text-xl pb-10">Add Member to Family</h1>

      <form onSubmit={handleSubmit}>
        <TextInput
          label="New Family Member Name"
          name="name"
          onChange={(ev) => setValue(ev.target.value)}
          value={value}
        />

        <Button type='submit'>Submit</Button>
      </form>
    </>
  )
}
