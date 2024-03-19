import ErrorModel from "../../error/ErrorModel";
import Bug from "../../model/Bug";
import apiV1BaseUrl from "../api";
import { BugsResponse } from "./bugResponse";

const getBugsByAssignedTeam = async (): Promise<BugsResponse> => {
  const res = await fetch(apiV1BaseUrl + "/bug/by-profile", {
    credentials: "include",
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const bugs = data as Bug[];
  return { bugs, error };
};

export default getBugsByAssignedTeam;
