import { useNavigate, useParams } from "react-router-dom";
import ErrorModel from "../../error/ErrorModel";
import TeamMemberForm from "../../component/teammember/TeamMemberForm";
import createTeamMember from "../../requests/teammember/createTeamMember";

const CreateTeamMember = () => {
  const navigate = useNavigate();
  const { id } = useParams();

  const handleDataFromChild = async (
    email: string,
    admin: boolean
  ): Promise<ErrorModel | null> => {
    const error = await createTeamMember(id ?? "", email, admin);
    if (error.errors) return error;
    else navigate(`/organisation/teams/${id}/members`);
    return null;
  };

  return (
    <TeamMemberForm
      teamId={id ?? ""}
      defaultEmail=""
      defaultAdmin={false}
      formTitle="Add Team Member"
      formButtonText="Add"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default CreateTeamMember;
