import { Routes, Route, Navigate } from "react-router-dom";
import Home from "../pages/Home";
import Login from "../pages/Login";
import Register from "../pages/Register";
import ListUrls from "../pages/ListUrls";

 
export default function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/urls" element={<ListUrls />} />

    
      
      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  );
}