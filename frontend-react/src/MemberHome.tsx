import AdminSection from './admin/AdminSection';
import { Link, useNavigate } from "react-router-dom";
import Button from './Button';
import { MouseEventHandler } from 'react';

interface MemberHomeType {
  familyMembership: any,
  doLogout: Function
}

export default function MemberHome({ familyMembership, doLogout }: MemberHomeType) {
  let navigate = useNavigate();

  const handleLogout: MouseEventHandler = (ev) => {
    ev.preventDefault()
    navigate('/');
    doLogout()
  }

  return <div className="flex flex-col p-10">
    <section>
      <h1>Welcome to Octopus Junior!</h1>
      <p>Here you will be able to explore create places where you can play and talk to your friends and family</p>
    </section>

    <section>
      <Button onClick={handleLogout}>logout</Button>
    </section>

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
