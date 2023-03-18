import { BrowserRouter, Routes, Route } from "react-router-dom";
import PersonPicker from "./PersonPicker";
import MeLayout from "./routes/MeLayout";
import useSessionStorage from "./util/useSessionStorage";

export default function UserLoggedInApp() {
  const [personId, setPersonId] = useSessionStorage(
    "personId",
    null
  );

  console.log("rendered LoggedInApp", personId);

  if (!personId) {
    return <PersonPicker setPersonId={setPersonId} />;
  }

  function doLogout() {
    setPersonId(null);
  }

  return (
    <BrowserRouter>
      <Routes>
        <Route path="*" element={<MeLayout doLogout={doLogout} />} />
      </Routes>
    </BrowserRouter>
  );
}
