import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MilestoneListPage from "./pages/MilestoneListPage";
import MilestoneDetailPage from "./pages/MilestoneDetailPage";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/milestones" element={<MilestoneListPage />} />
        <Route path="/milestones/:id" element={<MilestoneDetailPage />} />
      </Routes>
    </Router>
  );
}

export default App;
