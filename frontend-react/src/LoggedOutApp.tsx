import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import Button from "./Button";
import PersonAuth from "./PersonAuth";

export default function LoggedOutApp() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/kidsauth/*" element={<PersonAuth />} />
        <Route path="*" element={<AuthButtons />} />
      </Routes>
    </BrowserRouter>
  );
}

function AuthButtons() {
  const here = encodeURIComponent(window.location.href);

  return (
    <div className="grid h-screen place-items-center">
      <div className="flex flex-col items-center">
        <img alt="octopus" width="300px" src="octopus1.png" />
        <img
          alt="octopus junior text"
          width="300px"
          src="octopus-junior-text.png"
        />

        <Link to={`/kidsauth/login?from=${here}`}>
          <Button color="red">kids login</Button>
        </Link>
        <a href={`/auth/login?from=${here}`}>
          <Button color="blue">parents login</Button>
        </a>
      </div>
    </div>
  );
}
