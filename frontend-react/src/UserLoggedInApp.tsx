import { Routes, Route } from "react-router-dom";
import ParentIndex from "./routes/parent";
import PersonLoggedInApp from "./PersonLoggedInApp";
import ParentAdminHeader from "./ParentAdminHeader";


export default function UserLoggedInApp() {
  return (
    <div>
      <ParentAdminHeader />
      <Routes>
        <Route path="/parent/*" element={<ParentIndex />} />
        <Route path="*" element={<PersonLoggedInApp />} />
      </Routes>
    </div>
  );
}
