import React from 'react'
import { Link } from "react-router-dom";
import Button from './Button.js';

export default function MemberPageNotFound() {
  return (
    <div>
      <h1>page not found</h1>
      <Link to="/">
        <Button>go back to safety</Button>
      </Link>
    </div>
  )
}
