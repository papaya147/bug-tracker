import React from "react";
import TeamMember from "../../model/TeamMembers";
import { useNavigate } from "react-router-dom";

interface Props {
  teamId: string;
  members: TeamMember[];
}

const TeamMemberTable: React.FC<Props> = ({ teamId, members }) => {
  const navigate = useNavigate();

  return (
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Email</th>
          <th>Admin</th>
        </tr>
      </thead>
      <tbody>
        {members.map((member) => {
          return (
            <tr
              key={member.id}
              onClick={() =>
                navigate(
                  `/organisation/teams/${teamId}/members/${member.id}/update`
                )
              }
            >
              <td>{member.name}</td>
              <td>{member.email}</td>
              <td>{member.admin ? "Yes" : "No"}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

export default TeamMemberTable;
