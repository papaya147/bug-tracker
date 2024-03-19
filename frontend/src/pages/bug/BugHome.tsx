import { Link } from "react-router-dom";
import Profile from "../../model/Profile";
import { useEffect, useState } from "react";
import Bug from "../../model/Bug";
import getBugsByAssignedTeam from "../../requests/bug/getBugsByAsignedTeam";
import BugTable from "../../component/bug/BugTable";

interface Props {
  profile: Profile | null;
}

const BugHome: React.FC<Props> = ({ profile }) => {
  const [bugs, setBugs] = useState<Bug[] | null>(null);

  useEffect(() => {
    getBugsByAssignedTeam().then((data) => {
      if (!data.error.errors) setBugs(data.bugs);
    });
  }, []);

  return (
    <div>
      <Link to="/bug/create">Create a Bug!</Link>
      {bugs && <BugTable bugs={bugs} />}
    </div>
  );
};

export default BugHome;
