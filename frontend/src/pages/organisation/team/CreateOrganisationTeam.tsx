import { useNavigate } from "react-router-dom";
import TeamForm from "../../../component/team/TeamForm";
import ErrorModel from "../../../error/ErrorModel";
import createTeam from "../../../requests/team/createTeam";

const CreateOrganisationTeam = () => {
  const navigate = useNavigate();

  const handleDataFromChild = async (
    name: string,
    description: string
  ): Promise<ErrorModel | null> => {
    const data = await createTeam(name, description);
    if (data.error.errors) return data.error;
    else navigate(-1);
    return null;
  };

  return (
    <TeamForm
      defaultName=""
      defaultDescription=""
      formTitle="Create Team"
      formButtonText="Create"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default CreateOrganisationTeam;
