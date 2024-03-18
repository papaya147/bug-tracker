import ErrorModel from "../../error/ErrorModel";
import { useNavigate } from "react-router-dom";
import createOrganisation from "../../requests/organisation/createOrganisation";
import OrganisationForm from "../../component/organisation/OrganisationForm";

const CreateOrganisation = () => {
  const navigate = useNavigate();

  const handleDataFromChild = async (
    name: string,
    description: string
  ): Promise<ErrorModel | null> => {
    const data = await createOrganisation(name, description);
    if (data.error.errors) return data.error;
    else navigate(-1);
    return null;
  };

  return (
    <OrganisationForm
      defaultName=""
      defaultDescription=""
      formTitle="Create Organisation"
      formButtonText="Create"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default CreateOrganisation;
