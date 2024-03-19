import { useNavigate } from "react-router-dom";
import Team from "../../model/Team";

interface Props {
  team: Team;
  showUpdateButton: boolean;
  viewMembersLink: string;
}

const TeamCard: React.FC<Props> = ({
  team,
  showUpdateButton,
  viewMembersLink,
}) => {
  const navigate = useNavigate();

  return (
    <div className="card">
      <h3>{team.name}</h3>
      <p>{team.description}</p>
      {team.created_at && (
        <p>
          Established {new Date(team.created_at * 1000).toLocaleDateString()}
        </p>
      )}
      <h4>Organisation: {team.organisation_name}</h4>
      <p>{team.organisation_description}</p>
      {!showUpdateButton && (
        <button onClick={() => navigate(`/teams/${team.id}/bugs`)}>
          View Bugs
        </button>
      )}
      <button onClick={() => navigate(viewMembersLink)}>View Members</button>
      {showUpdateButton && (
        <button
          onClick={() => navigate(`/organisation/teams/${team.id}/update`)}
        >
          Update details
        </button>
      )}
    </div>
  );
};

export default TeamCard;
