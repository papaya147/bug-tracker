import Team from "../../model/Team";
import TeamCard from "./TeamCard";

interface Props {
  teams: Team[];
}

const TeamList: React.FC<Props> = ({ teams }) => {
  return (
    <div className="card-list">
      {teams.map((team) => {
        return <TeamCard team={team} key={team.id} />;
      })}
    </div>
  );
};

export default TeamList;
