import { Routes, Route } from "react-router-dom";
import PageNotFound from "../PageNotFound";
import AdminCreateSpace from "./AdminCreateSpace";

export default function AdminLayout() {
  return (
    <Routes>
      <Route path="create-space" element={<AdminCreateSpace />} />
      <Route path="*" element={<PageNotFound />} />
    </Routes>
  );
}
