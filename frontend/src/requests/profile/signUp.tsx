import ErrorModel from "../../error/ErrorModel";
import Profile from "../../model/Profile";
import apiV1BaseUrl from "../api";
import ProfileResponse from "./profileResponse";

const signUp = async (
  name: string,
  email: string,
  password: string
): Promise<ProfileResponse> => {
  const res = await fetch(apiV1BaseUrl + "/profile", {
    method: "POST",
    body: JSON.stringify({ name, email, password }),
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const profile = data as Profile;
  return { profile, error };
};

export default signUp;
