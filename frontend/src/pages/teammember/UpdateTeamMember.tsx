import { useNavigate, useParams } from "react-router-dom";
import getTeamMembers from "../../requests/teammember/getTeamMembers";
import { useEffect, useState } from "react";
import TeamMembers from "../../model/TeamMembers";
import ErrorModel from "../../error/ErrorModel";
import updateTeamMember from "../../requests/teammember/updateTeamMember";
import TeamMemberForm from "../../component/teammember/TeamMemberForm";

interface Props {
  isOrganisationChild: boolean;
}

const UpdateTeamMember: React.FC<Props> = ({ isOrganisationChild }) => {
  const [member, setMember] = useState<TeamMembers | null>(null);
  const { id, profile_id } = useParams();
  const navigate = useNavigate();

  useEffect(() => {
    getTeamMembers(id ?? "").then((data) => {
      if (!data.error.errors) {
        data.members.map((member) => {
          if (member.id === profile_id) setMember(member);
          return null;
        });
      } else navigate(-1);
    });
  }, [id, navigate, profile_id]);

  const handleDataFromChild = async (
    email: string,
    admin: boolean
  ): Promise<ErrorModel | null> => {
    const error = await updateTeamMember(id ?? "", profile_id ?? "", admin);
    if (error?.errors) return error;
    else navigate(-1);
    return null;
  };

  return (
    <TeamMemberForm
      teamId={id ?? ""}
      defaultEmail={member?.email}
      defaultAdmin={member?.admin}
      formTitle="Update Team Member"
      formButtonText="Update"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default UpdateTeamMember;
