import React, {useState} from 'react'
import { Routes, Route, Link } from "react-router-dom";
import { useMutation } from '@apollo/client';
import PageNotFound from '../PageNotFound';
import Button from '../Button';
import { CREATE_NEW_FAMILY_MEMBER } from '../queries.js';

export default function AdminLayout() {
  return (
    <Routes>
      <Route path="add-family-member" element={<AdminAddFamilyMember/>}/>
      <Route path="*" element={<PageNotFound />} />
    </Routes>
  )
}

function AdminAddFamilyMember() {
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

function TextInput({ label, name, onChange, value }) {
  return (
    <>
      <label className="form-label inline-block mb-2 text-gray-700">
        {label}
      </label>
      <input
        name={name}
        type="text"
        className={`
            form-control
            block
            w-full
            px-3
            py-1.5
            text-base
            font-normal
            text-gray-700
            bg-white bg-clip-padding
            border border-solid border-gray-300
            rounded
            transition
            ease-in-out
            m-0
            focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none
        `}
        onChange={onChange}
        value={value}
        placeholder={`${label}...`}
      />
    </>
  )
}
