import { useNavigate } from "react-router-dom";
import OrganisationForm from "../../component/organisation/OrganisationForm";
import ErrorModel from "../../error/ErrorModel";
import updateOrganisation from "../../requests/organisation/updateOrganisation";
import { useEffect, useState } from "react";
import getOrganisation from "../../requests/organisation/getOrganisation";
import Organisation from "../../model/Organisation";

const UpdateOrganisation = () => {
  const navigate = useNavigate();
  const [organisation, setOrganisation] = useState<Organisation | null>(null);

  useEffect(() => {
    getOrganisation().then((data) => {
      if (!data.error.errors) setOrganisation(data.organisation);
      else navigate("/organisation");
      console.log(data.organisation.name);
    });
  }, [navigate]);

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
      defaultName={organisation?.name}
      defaultDescription={organisation?.description}
      formTitle="Update Organisation Details"
      formButtonText="Update"
      sendDataToParent={handleDataFromChild}
    />
  );
};

export default UpdateOrganisation;
