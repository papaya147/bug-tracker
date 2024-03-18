import { useEffect, useState } from "react";
import ErrorModel from "../../error/ErrorModel";
import ServerError from "../../error/ServerError";
import { useNavigate } from "react-router-dom";
import ErrorList from "../../error/ErrorList";

interface Props {
  teamId: string;
  formTitle: string;
  formButtonText: string;
  defaultEmail: string | undefined;
  defaultAdmin: boolean | undefined;
  sendDataToParent: (
    email: string,
    admin: boolean
  ) => Promise<ErrorModel | null>;
}

const TeamMemberForm: React.FC<Props> = ({
  teamId,
  formTitle,
  formButtonText,
  defaultEmail,
  defaultAdmin,
  sendDataToParent,
}) => {
  const [email, setEmail] = useState("");
  const [admin, setAdmin] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<ErrorModel | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    if (defaultEmail) setEmail(defaultEmail);
    if (defaultAdmin) setAdmin(defaultAdmin);
  }, [defaultAdmin, defaultEmail]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    sendDataToParent(email, admin)
      .then((error) => {
        setError(error);
        setIsLoading(false);
      })
      .catch((e) => {
        setError(ServerError);
        setIsLoading(false);
      });
  };

  return (
    <div className="form">
      <h2>{formTitle}</h2>
      <form onSubmit={handleSubmit}>
        <label>Email</label>
        <input
          type="text"
          required
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <label>
          Admin
          <input
            type="checkbox"
            checked={admin}
            onChange={(e) => setAdmin(e.target.checked)}
          />
        </label>
        {!isLoading && <button type="submit">{formButtonText}</button>}
        {isLoading && <button disabled>Loading...</button>}
        {error && <ErrorList messages={error} />}
      </form>
      <button onClick={() => navigate(-1)}>Go Back</button>
    </div>
  );
};

export default TeamMemberForm;
