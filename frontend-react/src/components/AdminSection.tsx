import React, { MouseEventHandler } from 'react';
import Button from '../Button';
import { Link } from "react-router-dom";

export default function AdminSection() {
  const hardLogout: MouseEventHandler = (ev) => {
    ev.preventDefault()
    const yes = window.confirm('Are you sure you want to logout?')
    if (yes) {
      sessionStorage.clear()
      window.location.assign('/auth/logout');
    }
  }

  return (  
    <div className="flex gap-x-1 justify-start" >
      <Link to="/admin/add-family-member">
        <Button color="blue">add family member</Button>
      </Link>
      <Link to="/admin/create-space">
        <Button color="blue">create space</Button>
      </Link>
      <Button onClick={hardLogout} color="red">sign out completely</Button>
    </div>
  )
}
