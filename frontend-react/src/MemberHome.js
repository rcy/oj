import React from 'react'
import AdminSection from './admin/AdminSection.js';
import { Link } from "react-router-dom";
import Button from './Button.js';

export default function MemberHome({ familyMembership }) {
  return <div className="flex flex-col gap-y-10">
    <section>
      {familyMembership.role === 'admin' && <AdminSection />}
    </section>

    <section>
      <Link to="/spaces">
        <Button>Spaces</Button>
      </Link>
    </section>
  </div>
}
