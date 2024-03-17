import ErrorModel from "../../error/ErrorModel";
import Profile from "../../model/Profile";
import apiV1BaseUrl from "../api";

type SignUpResponse = {
  profile: Profile;
  error: ErrorModel;
};

const signUp = async (
  name: string,
  email: string,
  password: string
): Promise<SignUpResponse> => {
  const res = await fetch(apiV1BaseUrl + "/profile", {
    method: "POST",
    body: JSON.stringify({ name, email, password }),
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

export default signUp;
