import ErrorModel from "../../error/ErrorModel";
import Team from "../../model/Team";

type TeamsResponse = {
  teams: Team[];
  error: ErrorModel;
};

type TeamResponse = {
  team: Team;
  error: ErrorModel;
};

export type { TeamsResponse, TeamResponse };
