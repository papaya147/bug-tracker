import ErrorModel from "../../error/ErrorModel";
import Profile from "../../model/Profile";
import apiV1BaseUrl from "../api";
import ProfileResponse from "./profileResponse";

const checkSession = async (): Promise<ProfileResponse> => {
  const res = await fetch(apiV1BaseUrl + "/profile", {
    method: "GET",
    credentials: "include",
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const profile = data as Profile;
  return { profile, error };
};

export default checkSession;
