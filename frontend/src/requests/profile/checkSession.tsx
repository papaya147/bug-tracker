import ErrorModel from "../../error/ErrorModel";
import Profile from "../../model/Profile";
import apiV1BaseUrl from "../api";

type SessionResponse = {
  profile: Profile;
  error: ErrorModel;
};

const checkSession = async (): Promise<SessionResponse> => {
  const res = await fetch(apiV1BaseUrl + "/profile", {
    method: "GET",
    credentials: "include",
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const profile = data as Profile;
  return {
    profile: profile,
    error: error,
  };
};

export default checkSession;
