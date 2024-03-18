import ErrorModel from "../../error/ErrorModel";
import Team from "../../model/Team";
import apiV1BaseUrl from "../api";
import { TeamsResponse } from "./teamResponse";

const getTeams = async (): Promise<TeamsResponse> => {
  const res = await fetch(apiV1BaseUrl + "/team", {
    credentials: "include",
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const teams = data as Team[];
  return { teams, error };
};

export default getTeams;
