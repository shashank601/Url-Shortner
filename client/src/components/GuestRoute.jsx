import { Navigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext.jsx";

function GuestRoute({ children }) {
  const { user } = useAuth();

  if (user) {
    return <Navigate to="/" replace />;
  }

  return children;
}

export default GuestRoute;