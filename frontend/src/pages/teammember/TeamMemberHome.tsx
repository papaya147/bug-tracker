import { useState, useEffect } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import TeamMember from "../../model/TeamMembers";
import getTeamMembers from "../../requests/teammember/getTeamMembers";
import TeamMemberTable from "../../component/teammember/TeamMemberTable";

const TeamMemberHome = () => {
  const navigate = useNavigate();
  const [members, setMembers] = useState<TeamMember[] | null>(null);
  const { id } = useParams();

  useEffect(() => {
    getTeamMembers(id ?? "").then((data) => {
      if (!data.error.errors) setMembers(data.members);
      else navigate("/organisation/teams");
    });
  }, [navigate, id]);

  return (
    <div>
      <Link to={`/organisation/teams/${id}/members/create`}>Add a Member!</Link>
      {members && <TeamMemberTable teamId={id ?? ""} members={members} />}
    </div>
  );
};

export default TeamMemberHome;
