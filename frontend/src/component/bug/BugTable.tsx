import { useNavigate } from "react-router-dom";
import Bug from "../../model/Bug";
import { useState } from "react";
import ErrorModel from "../../error/ErrorModel";
import ErrorList from "../../error/ErrorList";

interface Props {
  bugs: Bug[];
  childOfAssigneeView: boolean;
  handleDelete?: (id: string) => Promise<ErrorModel | null>;
}

const BugTable: React.FC<Props> = ({
  bugs,
  childOfAssigneeView,
  handleDelete,
}) => {
  const navigate = useNavigate();
  const [error, setError] = useState<ErrorModel | null>(null);

  const del = (id: string) => {
    if (handleDelete)
      handleDelete(id).then((data) => {
        if (data?.errors) setError(data);
      });
  };

  return (
    <div>
      {error && <ErrorList messages={error} />}
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Description</th>
            <th>Priority</th>
            <th>Status</th>
            {!childOfAssigneeView && <th>Resolve</th>}
            {childOfAssigneeView && <th>Edit</th>}
            {childOfAssigneeView && <th>Delete</th>}
          </tr>
        </thead>
        <tbody>
          {bugs.map((bug) => {
            return (
              <tr key={bug.id}>
                <td>{bug.name}</td>
                <td>{bug.description}</td>
                <td>{bug.priority}</td>
                <td>{bug.status}</td>
                {!childOfAssigneeView && (
                  <td>
                    <button>Resolve</button>
                  </td>
                )}
                {childOfAssigneeView && (
                  <td>
                    <button onClick={() => navigate(`/bug/${bug.id}/update`)}>
                      Edit
                    </button>
                  </td>
                )}
                {childOfAssigneeView && (
                  <td>
                    <button onClick={() => del(bug.id)}>Delete</button>
                  </td>
                )}
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};

export default BugTable;
