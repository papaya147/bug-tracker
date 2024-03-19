import ErrorModel from "../../error/ErrorModel";
import Bug from "../../model/Bug";

type BugResponse = {
  bug: Bug;
  error: ErrorModel;
};

type BugsResponse = {
  bugs: Bug[];
  error: ErrorModel;
};

export type { BugResponse, BugsResponse };
