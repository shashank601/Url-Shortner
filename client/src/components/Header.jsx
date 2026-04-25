import { Link } from "react-router-dom";
import { useLocation } from "react-router-dom";
import { useState } from "react";

export default function Header() {
    const location = useLocation();
    

    return (
        <div className="sticky top-0 bg-zinc-950 p-4 flex justify-between font-mono items-center">
            <h1 className="text-2xl font-bold text-white">URL Shortner</h1>
            <nav className="flex justify-end gap-4 items-center">
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