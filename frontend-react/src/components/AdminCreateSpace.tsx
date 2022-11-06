import { FormEventHandler, useState } from 'react';
import { useCreateSpaceMutation } from '../generated-types';
import TextInput from './TextInput';
import Button from '../Button';

export default function AdminAddSpace() {
  const [value, setValue] = useState('')
  const [createSpace] = useCreateSpaceMutation();

  const handleSubmit: FormEventHandler = async (ev) => {
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
          onChange={(ev: React.ChangeEvent<HTMLInputElement>) => setValue(ev.target.value)}
          value={value}
        />

        <Button color="blue" type='submit'>Submit</Button>
      </form>
    </>
  )
}
