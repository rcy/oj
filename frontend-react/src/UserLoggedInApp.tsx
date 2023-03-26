import { Routes, Route } from "react-router-dom";
import MeLayout from "./routes/MeLayout";
import PersonPicker from "./PersonPicker";
import { logout } from "./util/logout";
import PersonLoggedInApp from "./PersonLoggedInApp";

export default function UserLoggedInApp() {
  return (
    <Routes>
      <Route path="/picker" element={<PersonPicker />} />
      <Route path="/oldme" element={<MeLayout doLogout={logout} />} />
      <Route path="*" element={<PersonLoggedInApp />} />
    </Routes>
  );
}
