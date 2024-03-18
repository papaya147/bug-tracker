import { useEffect, useState } from "react";
import Team from "../../model/Team";
import getTeams from "../../requests/team/getTeams";
import TeamList from "../../component/team/TeamList";

const TeamHome = () => {
  const [teams, setTeams] = useState<Team[] | null>(null);

  useEffect(() => {
    getTeams().then((data) => {
      console.log(data);
      if (!data.error.errors) setTeams(data.teams);
    });
  }, []);

  return (
    <div className="home-view">
      {teams && <TeamList teams={teams} isOrganisationChild={false} />}
    </div>
  );
};

export default TeamHome;
