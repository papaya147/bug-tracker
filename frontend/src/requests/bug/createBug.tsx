import { BugResponse } from "./bugResponse";
import apiV1BaseUrl from "../api";
import ErrorModel from "../../error/ErrorModel";
import Bug from "../../model/Bug";

const createBug = async (
  name: string,
  description: string,
  assignedTeam: string,
  assigneeTeam: string,
  priority: string
): Promise<BugResponse> => {
  const res = await fetch(apiV1BaseUrl + "/bug", {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({
      name,
      description,
      assigned_team: assignedTeam,
      assignee_team: assigneeTeam,
      priority,
    }),
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const bug = data as Bug;
  return { bug, error };
};

export default createBug;
