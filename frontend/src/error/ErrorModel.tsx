import ErrorDetail from "./ErrorDetail";

type ErrorModel = {
  detail: string;
  errors: ErrorDetail[];
  instance?: string;
  status: number;
  title: string;
  type?: string;
};

export default ErrorModel;
