import { useState } from "react";
import ErrorModel from "../../error/ErrorModel";
import { Link } from "react-router-dom";
import ErrorList from "../../error/ErrorList";
import ServerError from "../../error/ServerError";

interface Props {
  formTitle: string;
  formButtonText: string;
  sendDataToParent: (
    name: string,
    description: string
  ) => Promise<ErrorModel | null>;
}

const OrganisationForm: React.FC<Props> = ({
  formTitle,
  formButtonText,
  sendDataToParent,
}) => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<ErrorModel | null>(null);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    sendDataToParent(name, description)
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
        <label>Name</label>
        <input
          type="text"
          required
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <label>Description</label>
        <textarea
          required
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        {!isLoading && <button type="submit">{formButtonText}</button>}
        {isLoading && <button disabled>Loading...</button>}
        {error && <ErrorList messages={error} />}
      </form>
      <Link to={"/organisation"}>Go Back</Link>
    </div>
  );
};

export default OrganisationForm;
