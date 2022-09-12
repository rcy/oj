import React from 'react';
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";

export default function LoggedInApp() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>} />
      </Routes>
    </BrowserRouter>
  )
}

function Home() {
  return <p>home</p>
}
