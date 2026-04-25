import { Routes, Route, Navigate } from "react-router-dom";
import Home from "../pages/Home";
import Login from "../pages/Login";
import Register from "../pages/Register";
import ListUrls from "../pages/ListUrls";
import GuestRoute from "../components/GuestRoute";
import RequireAuth from "../components/RequireAuth";
import AppLayout from "../layout/AppLayout";

export default function AppRoutes() {
  return (
    <Routes>
      <Route
        path="/login"
        element={
          <GuestRoute>
            <Login />
          </GuestRoute>
        }
      />
      <Route
        path="/register"
        element={
          <GuestRoute>
            <Register />
          </GuestRoute>
        }
      />
      <Route element={<AppLayout />}>
        <Route path="/" element={<Home />} />
        <Route path="/urls" element={<ListUrls />} />
      </Route>

      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  );
}
