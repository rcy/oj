import React from 'react'
import AdminSection from './admin/AdminSection.js';

export default function MemberHome({ familyMembership }) {
  return <div>
    {familyMembership.role === 'admin' && <AdminSection />}
  </div>
}
