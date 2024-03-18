import { useEffect, useState } from "react";
import Team from "../../model/Team";
import getTeams from "../../requests/team/getTeams";
import TeamList from "../../component/team/TeamList";
import { Link } from "react-router-dom";

const TeamHome = () => {
  const [teams, setTeams] = useState<Team[] | null>(null);

  useEffect(() => {
    getTeams().then((data) => {
      console.log(data);
      if (!data.error.errors)
        setTeams(data.teams.length === 0 ? null : data.teams);
    });
  }, []);

  return (
    <div className="home-view">
      {teams && <TeamList teams={teams} isOrganisationChild={false} />}
      {!teams && (
        <div>
          <h2>You don't seem to be a part of any teams yet :(</h2>
          <p>
            You can ask people to add you to theirs or{" "}
            <Link to="/organisation">click here to create one.</Link>
          </p>
        </div>
      )}
    </div>
  );
};

export default TeamHome;
