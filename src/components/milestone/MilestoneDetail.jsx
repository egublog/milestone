import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const MilestoneDetail = () => {
  const [milestone, setMilestone] = useState({});
  const milestoneId = useParams().id;

  useEffect(() => {
    const fetchMilestone = async () => {
      try {
        const response = await fetch(
          `http://localhost:8080/api/milestones/${milestoneId}`
        );
        const data = await response.json();
        setMilestone(data);
      } catch (error) {
        console.error(error);
      }
    };
    fetchMilestone();
  }, [milestoneId]);

  return (
    <div>
      <h1>Milestone Detail</h1>
      <p>{milestone.title}</p>
      <p>{milestone.description}</p>
    </div>
  );
};

export default MilestoneDetail;
