import Team from "../../model/Team";
import TeamCard from "./TeamCard";

interface Props {
  teams: Team[];
  isOrganisationChild: boolean;
}

const TeamList: React.FC<Props> = ({ teams, isOrganisationChild }) => {
  return (
    <div className="card-list">
      {teams.map((team) => {
        return (
          <TeamCard
            team={team}
            showUpdateButton={isOrganisationChild}
            viewMembersLink={
              isOrganisationChild
                ? `/organisation/teams/${team.id}/members`
                : `/teams/${team.id}/members`
            }
            key={team.id}
          />
        );
      })}
    </div>
  );
};

export default TeamList;
