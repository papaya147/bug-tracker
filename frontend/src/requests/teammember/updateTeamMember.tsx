import ErrorModel from "../../error/ErrorModel";
import apiV1BaseUrl from "../api";

const updateTeamMember = async (
  team_id: string,
  profile_id: string,
  admin: boolean
): Promise<ErrorModel> => {
  const res = await fetch(apiV1BaseUrl + `/team-member`, {
    method: "PUT",
    credentials: "include",
    body: JSON.stringify({ team_id, profile_id, admin }),
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  const error = data as ErrorModel;
  return error;
};

export default updateTeamMember;
