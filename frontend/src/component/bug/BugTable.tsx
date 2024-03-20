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
      {bugs.length === 0 && <h3>No bugs to be found here :)</h3>}
      {bugs.length > 0 && (
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Description</th>
              <th>Priority</th>
              <th>Status</th>
              {childOfAssigneeView && <th>Completed</th>}
              {childOfAssigneeView && <th>Remark</th>}
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
                  {childOfAssigneeView && (
                    <td>{bug.completed ? "True" : "False"}</td>
                  )}
                  {childOfAssigneeView && <td>{bug.remarks}</td>}
                  {!childOfAssigneeView && (
                    <td>
                      <button
                        onClick={() => navigate(`/bug/${bug.id}/resolve`)}
                      >
                        Resolve
                      </button>
                    </td>
                  )}
                  {childOfAssigneeView &&
                    (!bug.completed ? (
                      <td>
                        <button
                          onClick={() => navigate(`/bug/${bug.id}/update`)}
                        >
                          Edit
                        </button>
                      </td>
                    ) : (
                      <td>-</td>
                    ))}
                  {childOfAssigneeView &&
                    (!bug.completed ? (
                      <td>
                        <button onClick={() => del(bug.id)}>Delete</button>
                      </td>
                    ) : (
                      <td>-</td>
                    ))}
                </tr>
              );
            })}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default BugTable;
