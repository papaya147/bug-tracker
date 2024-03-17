import Navbar from "./Navbar";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  useNavigate,
} from "react-router-dom";
import NotFound from "./NotFound";
import Home from "./Home";
import Login from "./Login";
import { useEffect, useState } from "react";
import Profile from "./model/Profile";
import checkSession from "./requests/profile/checkSession";

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

  useEffect(() => {
    checkSession().then((data) => {
      if (data.error.errors) navigate("/login");
      else {
        setProfile(data.profile);
        setLoggedIn(true);
      }
    });
  }, [navigate]);

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
          <Route path="*" element={<NotFound />} />
        </Routes>
      </div>
    </div>
  );
};

export default App;
