import ErrorModel from "./ErrorModel";

const ServerError: ErrorModel = {
  status: 500,
  title: "Server Unavailable",
  detail: "Server is unavailable.",
  errors: [{ message: "server is down/unavailable" }],
};

export default ServerError;
