import React from 'react';
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import FamilyLanding from './FamilyLanding.js'
import MemberHome from './MemberHome.js'

export default function LoggedInApp() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<FamilyLanding />} />
        <Route path="/member" element={<MemberHome />} />
      </Routes>
    </BrowserRouter>
  )
}
