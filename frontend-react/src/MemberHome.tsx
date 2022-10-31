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

  return <div className="flex flex-col">
    <div className='text-6xl'>
      Hi, {familyMembership?.person?.name}!
    </div>

    <section className='mt-5'>
      <h1>Welcome to Octopus Junior!</h1>
      <p>Here you will be able to explore and create places where you can play and talk to your friends and family</p>
    </section>

    <Link className="text-orange-500" to="/me/pic">
      <Button>select profile picture</Button>
    </Link>

    <section>
      <Button color="red" onClick={handleLogout}>logout</Button>
    </section>

    <section className='mt-10'>
      {familyMembership.role === 'admin' && <AdminSection />}
    </section>
  </div>
}
