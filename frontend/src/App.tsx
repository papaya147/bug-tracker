import Navbar from "./Navbar";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  useNavigate,
  useLocation,
} from "react-router-dom";
import NotFound from "./NotFound";
import Login from "./pages/profile/Login";
import { useEffect, useState } from "react";
import Profile from "./model/Profile";
import checkSession from "./requests/profile/checkSession";
import SignUp from "./pages/profile/SignUp";
import OrganisationHome from "./pages/organisation/OrganisationHome";
import CreateOrganisation from "./pages/organisation/CreateOrganisation";
import UpdateOrganisation from "./pages/organisation/UpdateOrganisation";
import TeamMemberHome from "./pages/teammember/TeamMemberHome";
import CreateTeamMember from "./pages/teammember/CreateTeamMember";
import UpdateTeamMember from "./pages/teammember/UpdateTeamMember";
import OrganisationTeamHome from "./pages/organisation/team/OrganisationTeamHome";
import CreateOrganisationTeam from "./pages/organisation/team/CreateOrganisationTeam";
import UpdateOrganisationTeam from "./pages/organisation/team/UpdateOrganisationTeam";
import TeamHome from "./pages/team/TeamHome";
import BugHome from "./pages/bug/BugHome";
import CreateBug from "./pages/bug/CreateBug";
import TeamBugHome from "./pages/team/TeamBugHome";
import UpdateBug from "./pages/bug/UpdateBug";

function App() {
  return (
    <Router>
      <Root />
    </Router>
  );
}

const Root = () => {
  const [loggedIn, setLoggedIn] = useState(false);
  const [profile, setProfile] = useState<Profile | null>(null);
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    if (location.pathname === "/login" || location.pathname === "/sign-up") {
      setLoggedIn(false);
      return;
    }

    checkSession()
      .then((data) => {
        if (data.error.errors) navigate("/login");
        else {
          setProfile(data.profile);
          setLoggedIn(true);
        }
      })
      .catch((e: Error) => {
        console.log(e);
      });
  }, [location.pathname, navigate]);

  return (
    <div className="App">
      {loggedIn && <Navbar />}
      <div className="content">
        <Routes>
          <Route
            path="/login"
            element={
              <Login setProfile={setProfile} setLoggedIn={setLoggedIn} />
            }
          />
          <Route path="/sign-up" element={<SignUp />} />
          <Route path="/organisation" element={<OrganisationHome />} />
          <Route path="/organisation/create" element={<CreateOrganisation />} />
          <Route path="/organisation/update" element={<UpdateOrganisation />} />
          <Route
            path="/organisation/teams"
            element={<OrganisationTeamHome />}
          />
          <Route
            path="/organisation/teams/create"
            element={<CreateOrganisationTeam />}
          />
          <Route
            path="/organisation/teams/:id/update"
            element={<UpdateOrganisationTeam />}
          />
          <Route
            path="/organisation/teams/:id/members"
            element={<TeamMemberHome />}
          />
          <Route
            path="/organisation/teams/:id/members/create"
            element={<CreateTeamMember />}
          />
          <Route
            path="/organisation/teams/:id/members/:profile_id/update"
            element={<UpdateTeamMember isOrganisationChild={true} />}
          />
          <Route path="/teams" element={<TeamHome />} />
          <Route path="/teams/:id/bugs" element={<TeamBugHome />} />
          <Route path="/teams/:id/members" element={<TeamMemberHome />} />
          <Route path="/" element={<BugHome profile={profile} />} />
          <Route path="/bug/create" element={<CreateBug />} />
          <Route path="/bug/:id/update" element={<UpdateBug />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </div>
    </div>
  );
};

export default App;
