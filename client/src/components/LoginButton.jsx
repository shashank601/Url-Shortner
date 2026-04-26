import React, {useState, useEffect} from 'react'
import { login, verify } from '../services/Auth.js';
import { useNavigate } from 'react-router-dom';
import { setToken } from '../utils/Token.js';
import { useAuth } from '../context/AuthContext.jsx';

export default function LoginButton({email, password}) {
	const [loading, setLoading] = useState(false);
	const navigate = useNavigate();
	const {setUser} = useAuth();

	const loginHandler = async () => {
		setLoading(true);
		try {
			const response = await login(email, password);
			setToken(response.token);
			
			// Get user data after setting token
			const userResponse = await verify();
			setUser(userResponse);
			
			navigate('/');
		} catch (err) {
			console.error(err);
		} finally {
			setLoading(false);
		}
	}


  return (
    <button onClick={loginHandler} disabled={loading} className="p-2 bg-zinc-900 text-white  ">login</button>
  )
}
