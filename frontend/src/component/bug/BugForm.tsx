import { useEffect, useState } from "react";
import ErrorModel from "../../error/ErrorModel";
import Organisation from "../../model/Organisation";
import ErrorList from "../../error/ErrorList";
import { useNavigate } from "react-router-dom";
import getAssignableOrganisations from "../../requests/bug/getAssignableOrganisations";
import getAssigneeTeams from "../../requests/bug/getAssigneeTeams";
import Team from "../../model/Team";

interface Props {
  formTitle: string;
  formButtonText: string;
  defaultName: string | undefined;
  defaultDescription: string | undefined;
  sendDataToParent: (
    name: string,
    description: string
  ) => Promise<ErrorModel | null>;
}

const BugForm: React.FC<Props> = ({ formTitle, formButtonText }) => {
  const [organisation, setOrganisation] = useState("");
  const [assignableOrganisations, setAssignableOrganisations] = useState<
    Organisation[] | null
  >(null);
  const [assigneeTeam, setAssigneeTeam] = useState("");
  const [organisationTeams, setOrganisationTeams] = useState<Team[] | null>(
    null
  );
  const [assignedTeam, setAssignedTeam] = useState("");
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [priority, setPriority] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<ErrorModel | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    getAssignableOrganisations().then((data) => {
      if (!data.error.errors) {
        setAssignableOrganisations(data.organisations);
        setOrganisation(data.organisations[0]?.id);
      } else setError(data.error);
    });
  }, []);

  useEffect(() => {
    if (!organisation) return;
    getAssigneeTeams(organisation).then((data) => {
      if (!data.error.errors) {
        setOrganisationTeams(data.teams);
        setAssigneeTeam(data.teams[0]?.id);
        setAssignedTeam(data.teams[0]?.id);
      } else setError(data.error);
    });
  }, [organisation]);

  const handleSubmit = (e: React.FormEvent) => {};

  return (
    <div className="form">
      <h2>{formTitle}</h2>
      <form onSubmit={handleSubmit}>
        <label>Organisation</label>
        <select
          value={organisation}
          onChange={(e) => setOrganisation(e.target.value)}
        >
          {assignableOrganisations &&
            assignableOrganisations.map((org) => {
              return (
                <option value={org.id} key={org.id}>
                  {org.name}
                </option>
              );
            })}
        </select>
        <label>Assignee Team</label>
        <select
          value={assigneeTeam}
          onChange={(e) => setAssigneeTeam(e.target.value)}
        >
          {organisationTeams &&
            organisationTeams.map((team) => {
              return (
                <option value={team.id} key={team.id}>
                  {team.name}
                </option>
              );
            })}
        </select>
        <label>Assigned Team</label>
        <select
          value={assignedTeam}
          onChange={(e) => setAssignedTeam(e.target.value)}
        >
          {organisationTeams &&
            organisationTeams.map((team) => {
              return (
                <option value={team.id} key={team.id}>
                  {team.name}
                </option>
              );
            })}
        </select>
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
      <button onClick={() => navigate(-1)}>Go Back</button>
    </div>
  );
};

export default BugForm;
