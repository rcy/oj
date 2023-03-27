import { Navigate, Route, Routes } from "react-router-dom";
import OnboardIntroduction from "./OnboardIntroduction";
import OnboardFamily from "./OnboardFamily";
import { Box } from "@chakra-ui/react";

export default function OnboardIndex() {
  return (
    <Box>
      <Routes>
        <Route path="/" element={<Navigate to="introduction" />} />
        <Route path="/introduction" element={<OnboardIntroduction />} />
        <Route path="/family/*" element={<OnboardFamily />} />
      </Routes>
    </Box>
  )
}
