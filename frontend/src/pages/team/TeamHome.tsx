import { useEffect, useState } from "react";
import Team from "../../model/Team";
import TeamList from "../../component/team/TeamList";
import getOrganisationTeams from "../../requests/team/getOrganisationTeams";
import { Link } from "react-router-dom";

const TeamHome = () => {
  const [teams, setTeams] = useState<Team[] | null>(null);

  useEffect(() => {
    getOrganisationTeams().then((data) => {
      console.log(data);
      if (!data.error.errors) setTeams(data.teams);
    });
  }, []);

  return (
    <div className="home-view">
      <Link to="/organisation/teams/create">Create a team!</Link>
      {teams && <TeamList teams={teams} />}
    </div>
  );
};

export default TeamHome;
