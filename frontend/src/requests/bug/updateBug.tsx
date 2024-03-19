import ErrorModel from "../../error/ErrorModel";
import Bug from "../../model/Bug";
import apiV1BaseUrl from "../api";
import { BugResponse } from "./bugResponse";

const updateBug = async (
  id: string,
  name: string,
  description: string,
  priority: string,
  status: string = "PENDING"
): Promise<BugResponse> => {
  const res = await fetch(apiV1BaseUrl + "/bug", {
    method: "PUT",
    credentials: "include",
    body: JSON.stringify({
      id,
      name,
      description,
      priority,
      status,
    }),
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const bug = data as Bug;
  return { bug, error };
};

export default updateBug;
