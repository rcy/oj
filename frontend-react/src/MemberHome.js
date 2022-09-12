import React from 'react'
import { CURRENT_MEMBERSHIP } from './queries.js'
import { useQuery } from '@apollo/client';
import { useNavigate } from 'react-router-dom';

export default function FamilyLanding() {
  const { loading, error, data } = useQuery(CURRENT_MEMBERSHIP)
  let navigate = useNavigate();

  if (loading) { return "loading" }

  return (
    <section>
      <h1>member home</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </section>
  )
}
