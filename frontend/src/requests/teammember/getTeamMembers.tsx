import ErrorModel from "../../error/ErrorModel";
import TeamMember from "../../model/TeamMembers";
import apiV1BaseUrl from "../api";
import TeamMembersResponse from "./teamMembersResponse";

const getTeamMembers = async (teamId: string): Promise<TeamMembersResponse> => {
  const res = await fetch(apiV1BaseUrl + `/team-member/${teamId}`, {
    method: "GET",
    credentials: "include",
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const members = data as TeamMember[];
  return { members, error };
};

export default getTeamMembers;
