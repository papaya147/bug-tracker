import Profile from "./model/Profile";

interface Props {
  profile: Profile | null; // Define the type of the profile prop
}

const Home: React.FC<Props> = ({ profile }) => {
  return <h1>Welcome{profile && " " + profile.name}!</h1>;
};

export default Home;
