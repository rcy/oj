import { Route, Routes } from "react-router-dom";
import PersonPage from "./PersonPage";

export default function PeopleIndexPage() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<h1>people index</h1>} />
        <Route path="/:id" element={<PersonPage />} />
      </Routes>
    </div>
  );
}
