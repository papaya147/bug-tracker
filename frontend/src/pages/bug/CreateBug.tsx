import { useNavigate } from "react-router-dom";
import BugForm from "../../component/bug/BugForm";
import ErrorModel from "../../error/ErrorModel";
import createBug from "../../requests/bug/createBug";

const CreateBug = () => {
  const navigate = useNavigate();

  const handleDataFromChild = async (
    name: string,
    description: string,
    assignedTeam: string,
    assigneeTeam: string,
    priority: string
  ): Promise<ErrorModel | null> => {
    const data = await createBug(
      name,
      description,
      assignedTeam,
      assigneeTeam,
      priority
    );
    if (data.error.errors) return data.error;
    else navigate(-1);
    return null;
  };

  return (
    <BugForm
      defaultName=""
      defaultDescription=""
      defaultAssignedTeam=""
      defaultAssigneeTeam=""
      defaultPriority=""
      formTitle="Create Bug"
      formButtonText="Create"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default CreateBug;
