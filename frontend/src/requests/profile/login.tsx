import ErrorModel from "../../error/ErrorModel";
import Profile from "../../model/Profile";

type LoginResponse = {
  profile: Profile;
  error: ErrorModel;
};

const login = async (
  email: string,
  password: string
): Promise<LoginResponse> => {
  const res = await fetch("http://localhost:4000/api/v1/profile/login", {
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
