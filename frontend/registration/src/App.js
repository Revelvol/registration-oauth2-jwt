import React, { useState } from "react";
import { googleLogout, useGoogleLogin } from "@react-oauth/google";
import axios from "axios"

function App() {
  const [user, setUser] = useState([]);
  const [profile, setProfile] = useState();


  const googleLogin = useGoogleLogin({
    flow: "auth-code",
    onSuccess: async (codeResponse) => {
      console.log(codeResponse);
      //todo this will proccess user and get its token from backend 
      const tokens = await axios.post("http://localhost:3001/api/auth/login", {
        code: codeResponse.code,
        channel: "google"
      });

      console.log(tokens);
    },
    onError: (errorResponse) => console.log(errorResponse),
  });

  // log out function to log the user out of google and set the profile array to null
  const logOut = () => {
    googleLogout();
    setProfile(null);
  };
  return (
    <div>
      <h2>React Google Login</h2>
      <br />
      <br />
      {profile ? (
        <div>
          <img src={profile.picture} alt="user image" />
          <h3>User Logged in</h3>
          <p>Name: {profile.name}</p>
          <p>Email Address: {profile.email}</p>
          <br />
          <br />
          <button onClick={logOut}>Log out</button>
        </div>
      ) : (
        <button onClick={googleLogin}>Sign in with Google 🚀 </button>
      )}
    </div>
  );
}
export default App;
