import ErrorModel from "../../error/ErrorModel";
import apiV1BaseUrl from "../api";

const createTeamMember = async (
  teamId: string,
  email: string,
  admin: boolean
): Promise<ErrorModel> => {
  const res = await fetch(apiV1BaseUrl + `/team-member`, {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({ team_id: teamId, email, admin }),
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  const error = data as ErrorModel;
  return error;
};

export default createTeamMember;
