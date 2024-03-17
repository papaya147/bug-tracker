import { Link, useNavigate } from "react-router-dom";
import Organisation from "../../model/Organisation";

interface Props {
  organisation: Organisation;
}

const OrganisationCard: React.FC<Props> = ({ organisation }) => {
  const navigate = useNavigate();

  return (
    <div className="card">
      <div className="container">
        <h3>{organisation.name}</h3>
        <p>{organisation.description}</p>
        <p>
          Established{" "}
          {new Date(organisation.created_at * 1000).toLocaleDateString()}
        </p>
        <button>View Teams</button>
        <button onClick={() => navigate("/organisation/update")}>
          Update details
        </button>
      </div>
    </div>
  );
};

export default OrganisationCard;
