import React from 'react'
import AdminSection from './admin/AdminSection';
import { Link } from "react-router-dom";
import Button from './Button';

type MemberHomeType = { familyMembership: any }

export default function MemberHome({ familyMembership }: MemberHomeType) {
  return <div className="flex flex-col gap-y-10">
    <section>
      {familyMembership.role === 'admin' && <AdminSection />}
    </section>

    <section>
      <Link to="/spaces">
        <Button color="blue">Spaces</Button>
      </Link>
    </section>
  </div>
}
