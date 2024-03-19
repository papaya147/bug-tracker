import { useNavigate, useParams } from "react-router-dom";
import ErrorModel from "../../error/ErrorModel";
import BugForm from "../../component/bug/BugForm";
import { useEffect, useState } from "react";
import Bug from "../../model/Bug";
import getBug from "../../requests/bug/getBug";
import updateBug from "../../requests/bug/updateBug";

const UpdateBug = () => {
  const navigate = useNavigate();
  const [bug, setBug] = useState<Bug | null>(null);
  const { id } = useParams();

  useEffect(() => {
    getBug(id ?? "").then((data) => {
      if (!data.error.errors) setBug(data.bug);
      else navigate(-1);
    });
  }, [id, navigate]);

  const handleDataFromChild = async (
    name: string,
    description: string,
    assignedTeam: string,
    assigneeTeam: string,
    priority: string
  ): Promise<ErrorModel | null> => {
    const data = await updateBug(id ?? "", name, description, priority);
    if (data.error.errors) return data.error;
    else navigate(-1);
    return null;
  };

  return (
    <BugForm
      defaultName={bug?.name}
      defaultDescription={bug?.description}
      editing={true}
      defaultPriority={bug?.priority}
      formTitle="Update Bug"
      formButtonText="Update"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default UpdateBug;
