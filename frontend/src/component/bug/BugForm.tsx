import { useEffect, useState } from "react";
import ErrorModel from "../../error/ErrorModel";
import Organisation from "../../model/Organisation";
import ErrorList from "../../error/ErrorList";
import { useNavigate } from "react-router-dom";
import getAssignableOrganisations from "../../requests/bug/getAssignableOrganisations";
import Team from "../../model/Team";
import ServerError from "../../error/ServerError";
import getAssignableTeams from "../../requests/bug/getAssignableTeams";
import getAssigneeTeams from "../../requests/bug/getAssigneeTeams";

interface Props {
  formTitle: string;
  formButtonText: string;
  defaultName: string | undefined;
  defaultDescription: string | undefined;
  editing: boolean | undefined;
  defaultPriority: string | undefined;
  sendDataToParent: (
    name: string,
    description: string,
    assignedTeam: string,
    assigneeTeam: string,
    priority: string
  ) => Promise<ErrorModel | null>;
}

const BugForm: React.FC<Props> = ({
  formTitle,
  formButtonText,
  defaultName,
  defaultDescription,
  editing,
  defaultPriority,
  sendDataToParent,
}) => {
  const [organisation, setOrganisation] = useState("");
  const [assignableOrganisations, setAssignableOrganisations] = useState<
    Organisation[] | null
  >(null);
  const [assigneeTeam, setAssigneeTeam] = useState("");
  const [assigneeTeams, setAssigneeTeams] = useState<Team[] | null>(null);
  const [assignedTeam, setAssignedTeam] = useState("");
  const [assignableTeams, setAssignableTeams] = useState<Team[] | null>(null);
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [priority, setPriority] = useState("LOW");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<ErrorModel | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    if (defaultName) setName(defaultName);
    if (defaultDescription) setDescription(defaultDescription);
    if (defaultPriority) setPriority(defaultPriority);
  }, [defaultDescription, defaultName, defaultPriority]);

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
        setAssigneeTeams(data.teams);
        setAssigneeTeam(data.teams[0]?.id);
      } else setError(data.error);
    });
    getAssignableTeams(organisation).then((data) => {
      if (!data.error.errors) {
        setAssignableTeams(data.teams);
        setAssignedTeam(data.teams[0]?.id);
      }
    });
  }, [organisation]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    sendDataToParent(name, description, assignedTeam, assigneeTeam, priority)
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
        {!editing && (
          <div>
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
          </div>
        )}
        {!editing && (
          <div>
            <label>Assignee Team</label>
            <select
              value={assigneeTeam}
              onChange={(e) => setAssigneeTeam(e.target.value)}
            >
              {assigneeTeams &&
                assigneeTeams.map((team) => {
                  return (
                    <option value={team.id} key={team.id}>
                      {team.name}
                    </option>
                  );
                })}
            </select>
          </div>
        )}
        {!editing && (
          <div>
            <label>Assigned Team</label>
            <select
              value={assignedTeam}
              onChange={(e) => setAssignedTeam(e.target.value)}
            >
              {assignableTeams &&
                assignableTeams.map((team) => {
                  return (
                    <option value={team.id} key={team.id}>
                      {team.name}
                    </option>
                  );
                })}
            </select>
          </div>
        )}
        <label>Name</label>
        <input
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
        <label>Priority</label>
        <select value={priority} onChange={(e) => setPriority(e.target.value)}>
          <option value="URGENT">Urgent</option>
          <option value="HIGH">High</option>
          <option value="LOW">Low</option>
        </select>
        {!isLoading && <button type="submit">{formButtonText}</button>}
        {isLoading && <button disabled>Loading...</button>}
        {error && <ErrorList messages={error} />}
      </form>
      <button onClick={() => navigate(-1)}>Go Back</button>
    </div>
  );
};

export default BugForm;
