import React from 'react';
import Button from './Button.js';

function addFamilyMember(ev) {
  alert('not implemented')
}

export default function AdminSection() {
  function hardLogout(ev) {
    ev.preventDefault()
    sessionStorage.clear()
    window.location = '/auth/logout'
  }

  return (  
    <div className="flex space-x-10" >
      <Button onClick={addFamilyMember}>add family member</Button>
      <Button onClick={hardLogout} color="red">sign out completely</Button>
    </div>
  )
}
