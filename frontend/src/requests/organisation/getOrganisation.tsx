import ErrorModel from "../../error/ErrorModel";
import Organisation from "../../model/Organisation";
import apiV1BaseUrl from "../api";
import OrganisationResponse from "./organisationResponse";

const getOrganisation = async (): Promise<OrganisationResponse> => {
  const res = await fetch(apiV1BaseUrl + "/organisation", {
    method: "GET",
    credentials: "include",
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const organisation = data as Organisation;
  return { organisation, error };
};

export default getOrganisation;
