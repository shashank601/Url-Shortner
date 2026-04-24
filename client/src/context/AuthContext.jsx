import { createContext, useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getToken, clearToken, setToken } from "../utils/Token.js";

import { verify } from "../services/Auth.js";

export const AuthContext = createContext(null);



export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();
  

  useEffect(() => {
    setToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImEiLCJpZCI6MywibmFtZSI6ImEifQ.a_o0JDvoOopMb9Aztx6mOsDLsSKOMyDQbQAknzWpuow");
    const token = getToken();
    
    if (token) {
      verify()
        .then(() => {
          setUser({ token });
        })
        .catch((err) => {
          clearToken();
        })
        .finally(() => {
          setLoading(false);
        });
    } else {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    const handleStorageChange = (e) => {
      if (e.key === "token") {
        const newToken = e.newValue;
        const currentToken = getToken();

        if (newToken !== currentToken) {
          clearToken();
          setUser(null);
          navigate("/login");
        }
      }
    };

    window.addEventListener("storage", handleStorageChange);

    return () => {
      window.removeEventListener("storage", handleStorageChange);
    };
  }, [navigate]);

  if (loading) {
    return null; 
  } else {
    return (
      <AuthContext.Provider value={{ user, setUser, loading }}>
        {children}
      </AuthContext.Provider>
    );
  }
};

export const useAuth = () => {
  const ctx = useContext(AuthContext);
  if (!ctx) {
    throw new Error('useAuth must be used inside AuthProvider');
  }
  return ctx;
};

