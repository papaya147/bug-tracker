import ErrorModel from "../../error/ErrorModel";
import Bug from "../../model/Bug";
import apiV1BaseUrl from "../api";
import { BugResponse } from "./bugResponse";

const closeBug = async (id: string, remarks: string): Promise<BugResponse> => {
  const res = await fetch(apiV1BaseUrl + "/bug/close", {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({ id, remarks }),
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const bug = data as Bug;
  return { bug, error };
};

export default closeBug;
