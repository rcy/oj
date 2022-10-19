import React from 'react';
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import FamilyLanding from './FamilyLanding'
import MemberLayout from './MemberLayout'
import useSessionStorage from './util/useSessionStorage'

export default function LoggedInApp() {
  const [familyMembershipId, setFamilyMembershipId] = useSessionStorage('familyMembershipId', null);

  console.log('rendered LoggedInApp', familyMembershipId)

  if (!familyMembershipId) {
    // rename to FamilyMemberPicker
    return <FamilyLanding setFamilyMembershipId={setFamilyMembershipId} />
  }

  function doLogout() {
    setFamilyMembershipId(null)
  }

  return (
    <BrowserRouter>
      <Routes>
        <Route path="*" element={<MemberLayout doLogout={doLogout}/>} />
      </Routes>
    </BrowserRouter>
  )
}
