import React from "react";
import "./PostBoard.css";
import "./Post.css";
import lahui from "../../assets/lahui.svg";
import { useNavigate } from "react-router-dom";
import PageHeader from "../../components/PageHeader"; // <<< นำเข้ามาใช้

const PostBoard: React.FC = () => {
  const navigate = useNavigate();

  return (
    <div className="bg-gray">
      <PageHeader title="บอร์ดประกาศงาน" /> {/* <<< ใช้ component ที่สร้างไว้ */}
      
      <div className="bg-post" onClick={() => navigate("/post-detail")}>
        <img src={lahui} alt="ล่าฮุย" className="post-image" />
      </div>
    </div>
  );
};
export default PostBoard;
