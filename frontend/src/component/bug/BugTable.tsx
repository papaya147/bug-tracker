import Bug from "../../model/Bug";

interface Props {
  bugs: Bug[];
}

const BugTable: React.FC<Props> = ({ bugs }) => {
  return (
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Email</th>
          <th>Priority</th>
          <th>Status</th>
          <th>Edit</th>
          <th>Delete</th>
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
              <td>
                <button>Edit</button>
              </td>
              <td>
                <button>Delete</button>
              </td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

export default BugTable;
