import React from 'react';
import Button from '../Button.js';
import { Link } from "react-router-dom";

export default function AdminSection() {
  function hardLogout(ev) {
    ev.preventDefault()
    const yes = window.confirm('Are you sure you want to logout?')
    if (yes) {
      sessionStorage.clear()
      window.location = '/auth/logout'
    }
  }

  return (  
    <div className="flex space-x-10" >
      <Link to="/admin/add-family-member">
        <Button>add family member</Button>
      </Link>
      <Button onClick={hardLogout} color="red">sign out completely</Button>
    </div>
  )
}
