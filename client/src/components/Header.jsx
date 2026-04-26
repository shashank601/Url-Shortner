import { Link } from "react-router-dom";
import { useLocation } from "react-router-dom";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { logout } from "../services/Auth";
import { useAuth } from "../context/AuthContext.jsx";

export default function Header() {
    const location = useLocation();
    const navigate = useNavigate();
    const { setUser } = useAuth();
    

    return (
        <div className="sticky top-0 bg-zinc-950 p-4 flex justify-between font-mono items-center">
            <h1 className="text-2xl font-bold text-white">URL Shortner</h1>
            <nav className="flex justify-end gap-4 items-center">
                <h1 className="text-blue-500 text-xl font-semibold cursor-pointer" onClick={() => {logout(); setUser(null); navigate('/login');}}>Logout</h1>
                {
                    location.pathname === '/' ? (
                        <Link to="/urls" className="text-blue-500 text-xl font-semibold">History</Link>
                        

                    ) : (
                        <Link to="/" className="text-blue-500 text-xl font-semibold">Home</Link>
                        
                    )
                }
            </nav>
        </div>
    );
}