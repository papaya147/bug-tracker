import ErrorModel from "../../error/ErrorModel";
import Organisation from "../../model/Organisation";
import apiV1BaseUrl from "../api";
import OrganisationResponse from "./organisationResponse";

const updateOrganisation = async (
  name: string,
  description: string
): Promise<OrganisationResponse> => {
  const res = await fetch(apiV1BaseUrl + "/organisation", {
    method: "PUT",
    credentials: "include",
    body: JSON.stringify({ name, description }),
    headers: { "Content-Type": "application/json" },
  });
  const data = await res.json();
  const error = data as ErrorModel;
  const organisation = data as Organisation;
  return { organisation, error };
};

export default updateOrganisation;
