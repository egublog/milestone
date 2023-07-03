import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MilestoneListPage from "./pages/MilestoneListPage";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<MilestoneListPage />} />
      </Routes>
    </Router>
  );
}

export default App;
