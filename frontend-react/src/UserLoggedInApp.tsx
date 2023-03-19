import { BrowserRouter, Routes, Route } from "react-router-dom";
import MeLayout from "./routes/MeLayout";
import PersonPicker from "./PersonPicker";
import { logout } from "./util/logout";

export default function UserLoggedInApp() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/picker" element={<PersonPicker />} />
        <Route path="*" element={<MeLayout doLogout={logout} />} />
      </Routes>
    </BrowserRouter>
  );
}
