import React from "react";
import { Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import Home from "./pages/Home/Home";
import JobPost from "./pages/JopPost/JobPost";
import Board from "./pages/Board/Board";
import JobDetail from "./pages/Board/JobDetail";


const App: React.FC = () => {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/board" element={<Board />} />
        <Route path="/post-job" element={<JobPost />} />
<Route path="/post-detail" element={<JobDetail />} />

      </Routes>
    </>
  );
};

export default App;
