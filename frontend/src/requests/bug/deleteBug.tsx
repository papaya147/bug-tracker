import ErrorModel from "../../error/ErrorModel";
import Bug from "../../model/Bug";
import apiV1BaseUrl from "../api";
import { BugResponse } from "./bugResponse";

const deleteBug = async (id: string): Promise<BugResponse> => {
  const res = await fetch(apiV1BaseUrl + `/bug/${id}`, {
    method: "DELETE",
    credentials: "include",
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const bug = data as Bug;
  return { bug, error };
};

export default deleteBug;
