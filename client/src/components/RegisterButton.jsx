import {login, register, verify } from '../services/Auth.js';
import {useState} from 'react'
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext.jsx';
import { setToken } from '../utils/Token.js';

export default function RegisterButton({email, password, username}) {
	const [loading, setLoading] = useState(false);
	const navigate = useNavigate();
	const {setUser} = useAuth();

	const registerHandler = async () => {
		setLoading(true);
		try {
			await register(username, email, password);

			const loginResponse = await login(email, password);
			setToken(loginResponse.token);

			const verifyResponse = await verify();
			setUser(verifyResponse);
            navigate('/');
		} catch (err) {
			console.error(err);
		} finally {
			setLoading(false);
		}
	}


  return (
    <button onClick={registerHandler} disabled={loading} className="p-2 bg-zinc-900 text-white  ">register</button>
  )
}
