import ErrorModel from "../../error/ErrorModel";
import Profile from "../../model/Profile";

type ProfileResponse = {
  profile: Profile;
  error: ErrorModel;
};

export default ProfileResponse;
