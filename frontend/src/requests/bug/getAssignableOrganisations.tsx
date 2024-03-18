import ErrorModel from "../../error/ErrorModel";
import Organisation from "../../model/Organisation";
import apiV1BaseUrl from "../api";

type GetAssignableOrganisationsResponse = {
  organisations: Organisation[];
  error: ErrorModel;
};

const getAssignableOrganisations =
  async (): Promise<GetAssignableOrganisationsResponse> => {
    const res = await fetch(apiV1BaseUrl + "/bug/organisations", {
      credentials: "include",
    });
    const data = await res.json();
    const error = data as ErrorModel;
    const organisations = data as Organisation[];
    return { organisations, error };
  };

export default getAssignableOrganisations;
