import { FormEventHandler, useState } from 'react';
import { useCreateNewFamilyMemberMutation } from '../generated-types';
import TextInput from './TextInput';
import Button from '../Button';

export default function AdminAddFamilyMember() {
  const [value, setValue] = useState('')
  const [addFamilyMember] = useCreateNewFamilyMemberMutation();

  const handleSubmit: FormEventHandler = async (ev) => {
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
          onChange={(ev: React.ChangeEvent<HTMLInputElement>) => setValue(ev.target.value)}
          value={value}
        />
        
        <Button color="green" type='submit'>Submit</Button>
      </form>
    </>
  )
}
