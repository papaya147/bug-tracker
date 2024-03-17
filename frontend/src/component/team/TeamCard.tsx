import { useNavigate } from "react-router-dom";
import Team from "../../model/Team";

interface Props {
  team: Team;
}

const TeamCard: React.FC<Props> = ({ team }) => {
  const navigate = useNavigate();

  return (
    <div className="card">
      <h3>{team.name}</h3>
      <p>{team.description}</p>
      <p>Established {new Date(team.created_at * 1000).toLocaleDateString()}</p>
      <button
        onClick={() => navigate(`/organisation/teams/${team.id}/members`)}
      >
        View Members
      </button>
      <button onClick={() => navigate(`/organisation/teams/${team.id}/update`)}>
        Update details
      </button>
    </div>
  );
};

export default TeamCard;
