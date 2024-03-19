import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Bug from "../../model/Bug";
import getBugsByAssigneeTeam from "../../requests/bug/getBugsByAssigneeTeam";
import BugTable from "../../component/bug/BugTable";
import ErrorModel from "../../error/ErrorModel";
import deleteBug from "../../requests/bug/deleteBug";

const TeamBugHome = () => {
  const { id } = useParams();
  const [bugs, setBugs] = useState<Bug[] | null>(null);
  const [refresh, setRefresh] = useState(false);

  useEffect(() => {
    getBugsByAssigneeTeam(id ?? "").then((data) => {
      if (!data.error.errors) setBugs(data.bugs);
    });
  }, [id, refresh]);

  const handleDelete = async (id: string): Promise<ErrorModel | null> => {
    const data = await deleteBug(id);
    if (data.error.errors) return data.error;
    setRefresh(!refresh);
    return null;
  };

  return (
    <div>
      {bugs && (
        <BugTable
          bugs={bugs}
          childOfAssigneeView={true}
          handleDelete={handleDelete}
        />
      )}
    </div>
  );
};

export default TeamBugHome;
