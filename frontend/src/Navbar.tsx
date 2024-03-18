import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav className="navbar">
      <h1>Buggy</h1>
      <div className="links">
        <Link to="/">Bugs</Link>
        <Link to="/organisation">Your Organisation</Link>
        <Link to="/teams">Your Teams</Link>
      </div>
    </nav>
  );
};

export default Navbar;
