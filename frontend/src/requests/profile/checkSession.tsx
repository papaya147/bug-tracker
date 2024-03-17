import ErrorModel from "../../error/ErrorModel";
import Profile from "../../model/Profile";

type SessionResponse = {
  profile: Profile;
  error: ErrorModel;
};

const checkSession = async (): Promise<SessionResponse> => {
  const res = await fetch("http://localhost:4000/api/v1/profile", {
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
