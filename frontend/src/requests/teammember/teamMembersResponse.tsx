import ErrorModel from "../../error/ErrorModel";
import TeamMembers from "../../model/TeamMembers";

type TeamMembersResponse = {
  members: TeamMembers[];
  error: ErrorModel;
};

export default TeamMembersResponse;
