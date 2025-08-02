import React from "react";
import { Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import Home from "./pages/Home/Home"; // ตามโครงสร้างจริง
import JobPost from "./pages/JopPost/JobPost";

const App: React.FC = () => {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/post-job" element={<JobPost />} />
      </Routes>
    </>
  );
};

export default App;
