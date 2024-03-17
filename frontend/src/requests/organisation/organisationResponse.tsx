import ErrorModel from "../../error/ErrorModel";
import Organisation from "../../model/Organisation";

type OrganisationResponse = {
  organisation: Organisation;
  error: ErrorModel;
};

export default OrganisationResponse;
