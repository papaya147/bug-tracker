import Navbar from "./Navbar";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  useNavigate,
  useLocation,
} from "react-router-dom";
import NotFound from "./NotFound";
import Home from "./pages/Home";
import Login from "./pages/Login";
import { useEffect, useState } from "react";
import Profile from "./model/Profile";
import checkSession from "./requests/profile/checkSession";
import SignUp from "./pages/SignUp";

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
          <Route path="/" element={<Home profile={profile} />} />
          <Route path="/sign-up" element={<SignUp />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </div>
    </div>
  );
};

export default App;
