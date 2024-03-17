import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav className="navbar">
      <h1>Buggy</h1>
      <div className="links">
        <Link to="/">Home</Link>
        <Link to="/organisation">Organisation</Link>
      </div>
    </nav>
  );
};

export default Navbar;
