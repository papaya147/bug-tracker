import { useState } from "react";
import ErrorModel from "../error/ErrorModel";
import ErrorList from "../error/ErrorList";
import { useNavigate } from "react-router-dom";
import signUp from "../requests/profile/signUp";
import ServerError from "../error/ServerError";

const SignUp = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<ErrorModel | null>(null);
  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    setIsLoading(true);

    if (password !== confirmPassword) {
      setError({
        status: 400,
        title: "Bad Request",
        detail: "Passwords do not match",
        errors: [{ message: "passwords do not match" }],
      });
      setIsLoading(false);
      return;
    }

    signUp(name, email, password)
      .then((data) => {
        if (data.error.errors) setError(data.error);
        else {
          alert(
            "Sign up successful! Verification email has been sent to your email."
          );
          navigate("/login");
        }
        setIsLoading(false);
      })
      .catch((e: Error) => {
        setError(ServerError);
        setIsLoading(false);
      });
  };

  return (
    <div className="form">
      <h2>Sign Up</h2>
      <form onSubmit={handleSubmit}>
        <label>Name</label>
        <input
          type="text"
          required
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
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
        <label>Confirm Password</label>
        <input
          type="password"
          required
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
        />
        {!isLoading && <button type="submit">Login</button>}
        {isLoading && <button disabled>Loading...</button>}
        {error && <ErrorList messages={error} />}
      </form>
    </div>
  );
};

export default SignUp;
