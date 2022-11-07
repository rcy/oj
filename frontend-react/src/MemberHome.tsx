import { Link, useNavigate } from "react-router-dom";
import Button from './Button';
import { MouseEventHandler } from 'react';
import FamilyIndex from "./routes/family";

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

  return (
    <div className="flex flex-col">
      <div className='text-6xl flex items-center justify-between pr-2 h-20 mr-2 my-2'>
        <div>
          Hi, {familyMembership?.person?.name}
        </div>
        <Button color="red" onClick={handleLogout}>logout</Button>
      </div>

      <section className='mt-1'>
        <h1 className="text-xl mb-2">Welcome to Octopus Junior!</h1>
        <p>Here you will be able to explore and create places where you can play and talk to your friends and family.</p>
      </section>

      <section className='mt-5 border border-solid-2 border-black bg-gray-100 p-2'>
        <Link className="text-blue-500" to="/me/pic">
          <div className="flex gap-2"><img width="20" src={familyMembership.person.avatarUrl} /> change my picture</div>
        </Link>
      </section>

      <section className='mt-5 border border-solid-2 border-black bg-gray-100 p-2'>
        <FamilyIndex />
      </section>
    </div >
  )
}
