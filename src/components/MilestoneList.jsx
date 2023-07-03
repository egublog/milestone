import React, { useEffect, useState } from "react";

const MilestoneList = () => {
  const [milestones, setMilestones] = useState([]);

  useEffect(() => {
    // APIリクエストを行う関数
    const fetchMilestones = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/milestones");
        const data = await response.json();
        setMilestones(data);
      } catch (error) {
        console.error(error);
      }
    };
    fetchMilestones();
  }, []);

  return (
    <div>
      <h1>Milestones List</h1>
      <ul>
        {milestones.map((milestone) => (
          <li key={milestone.ID}>{milestone.Title}</li>
        ))}
      </ul>
    </div>
  );
};

export default MilestoneList;