import ErrorModel from "../../error/ErrorModel";
import Profile from "../../model/Profile";
import apiV1BaseUrl from "../api";

type LoginResponse = {
  profile: Profile;
  error: ErrorModel;
};

const login = async (
  email: string,
  password: string
): Promise<LoginResponse> => {
  const res = await fetch(apiV1BaseUrl + "/profile/login", {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({ email, password }),
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const profile = data as Profile;
  return {
    profile: profile,
    error: error,
  };
};

export default login;
