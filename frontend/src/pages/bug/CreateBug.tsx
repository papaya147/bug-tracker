import { useNavigate } from "react-router-dom";
import BugForm from "../../component/bug/BugForm";
import ErrorModel from "../../error/ErrorModel";

const CreateBug = () => {
  const navigate = useNavigate();

  const handleDataFromChild = async (
    name: string,
    description: string
  ): Promise<ErrorModel | null> => {
    console.log(name);
    return null;
  };

  return (
    <BugForm
      defaultName=""
      defaultDescription=""
      formTitle="Create Bug"
      formButtonText="Create"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default CreateBug;
