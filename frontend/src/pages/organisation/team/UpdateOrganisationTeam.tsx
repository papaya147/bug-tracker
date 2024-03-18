import { useNavigate, useParams } from "react-router-dom";
import TeamForm from "../../../component/team/TeamForm";
import ErrorModel from "../../../error/ErrorModel";
import updateTeam from "../../../requests/team/updateTeam";
import { useEffect, useState } from "react";
import Team from "../../../model/Team";
import getOrganisationTeams from "../../../requests/team/getOrganisationTeams";

const UpdateOrganisationTeam = () => {
  const navigate = useNavigate();
  const { id } = useParams();
  const [team, setTeam] = useState<Team | null>(null);

  useEffect(() => {
    getOrganisationTeams().then((data) => {
      if (!data.error.errors) {
        data.teams.map((team) => {
          if (team.id === id) setTeam(team);
          return null;
        });
      } else navigate(-1);
    });
  }, [id, navigate]);

  const handleDataFromChild = async (
    name: string,
    description: string
  ): Promise<ErrorModel | null> => {
    const data = await updateTeam(id ?? "", name, description);
    if (data.error.errors) return data.error;
    else navigate(-1);
    return null;
  };

  return (
    <TeamForm
      defaultName={team?.name}
      defaultDescription={team?.description}
      formTitle="Update Team"
      formButtonText="Update"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default UpdateOrganisationTeam;
