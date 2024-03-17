import { useEffect, useState } from "react";
import Organisation from "../../model/Organisation";
import getOrganisation from "../../requests/organisation/getOrganisation";
import { Link } from "react-router-dom";
import OrganisationCard from "../../component/organisation/OrganisationCard";

const OrganisationHome = () => {
  const [organisation, setOrganisation] = useState<Organisation | null>(null);

  useEffect(() => {
    getOrganisation().then((data) => {
      if (!data.error.errors) setOrganisation(data.organisation);
    });
  }, []);

  return (
    <div className="home-view">
      {organisation && <OrganisationCard organisation={organisation} />}
      {!organisation && (
        <Link to="/organisation/create">Create your organisation!</Link>
      )}
    </div>
  );
};

export default OrganisationHome;
