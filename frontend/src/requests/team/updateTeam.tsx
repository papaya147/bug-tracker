import ErrorModel from "../../error/ErrorModel";
import Team from "../../model/Team";
import apiV1BaseUrl from "../api";
import { TeamResponse } from "./teamResponse";

const updateTeam = async (
  id: string,
  name: string,
  description: string
): Promise<TeamResponse> => {
  const res = await fetch(apiV1BaseUrl + `/organisation/team/${id}`, {
    method: "PUT",
    credentials: "include",
    body: JSON.stringify({ name, description }),
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const team = data as Team;
  return { team, error };
};

export default updateTeam;
