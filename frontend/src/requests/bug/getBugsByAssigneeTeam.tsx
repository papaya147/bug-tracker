import ErrorModel from "../../error/ErrorModel";
import Bug from "../../model/Bug";
import apiV1BaseUrl from "../api";
import { BugsResponse } from "./bugResponse";

const getBugsByAssigneeTeam = async (
  team_id: string
): Promise<BugsResponse> => {
  const res = await fetch(apiV1BaseUrl + `/bug/by-assignee-team/${team_id}`, {
    credentials: "include",
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const bugs = data as Bug[];
  return { bugs, error };
};

export default getBugsByAssigneeTeam;
