import React, { useEffect, useState } from "react";
import ErrorList from "./error/ErrorList";
import ErrorModel from "./error/ErrorModel";
import Profile from "./model/Profile";
import { useNavigate } from "react-router-dom";
import checkSession from "./requests/profile/checkSession";
import login from "./requests/profile/login";

type SetProfileType = React.Dispatch<React.SetStateAction<Profile | null>>;
type SetLoggedInType = React.Dispatch<React.SetStateAction<boolean>>;

interface Props {
  setProfile: SetProfileType;
  setLoggedIn: SetLoggedInType;
}

const Login: React.FC<Props> = ({ setProfile, setLoggedIn }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<ErrorModel | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    checkSession().then((data) => {
      if (!data.error.errors) {
        setProfile(data.profile);
        setLoggedIn(true);
        navigate("/");
      }
    });
  }, [navigate, setLoggedIn, setProfile]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    setIsLoading(true);

    login(email, password).then((data) => {
      if (data.error.errors) setError(data.error);
      else {
        setProfile(data.profile);
        navigate("/");
      }
      setIsLoading(false);
    });
  };

  return (
    <div className="login">
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <label>Email</label>
        <input
          type="text"
          required
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <label>Password</label>
        <input
          type="password"
          required
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        {!isLoading && <button>Login</button>}
        {isLoading && <button disabled>Loading...</button>}
        {error && <ErrorList messages={error} />}
      </form>
    </div>
  );
};

export default Login;
