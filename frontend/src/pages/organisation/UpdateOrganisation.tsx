import { useNavigate } from "react-router-dom";
import OrganisationForm from "../../component/organisation/OrganisationForm";
import ErrorModel from "../../error/ErrorModel";
import updateOrganisation from "../../requests/organisation/updateOrganisation";

const UpdateOrganisation = () => {
  const navigate = useNavigate();

  const handleDataFromChild = async (
    name: string,
    description: string
  ): Promise<ErrorModel | null> => {
    const data = await updateOrganisation(name, description);
    if (data.error.errors) return data.error;
    else navigate("/organisation");
    return null;
  };

  return (
    <OrganisationForm
      formTitle="Update Organisation Details"
      formButtonText="Update"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default UpdateOrganisation;
