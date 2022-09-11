import {useState, useCallback, useEffect} from "react"
import { setCookie, getCookie, deleteCookie } from "../middleware/cookieUtils"


export const useAuth = () => {

	
	const [token, setToken] = useState<string | null>(null)
	const [userId, setUserId] = useState<number | null>(null)
	const [ready, setReady] = useState<boolean>(false)


	 const login = useCallback((jwtToken: string, id: number): null => {
		setToken(jwtToken)
		setUserId(id)

		setCookie('userData', JSON.stringify({
			userId: id, token: jwtToken
		}))
		return null
	}, [])


	const logout = useCallback(() => {
		setToken(null)
		setUserId(null)
		deleteCookie('userData')

	}, [])

	useEffect(() => {
        const userData = getCookie('userData')
		if(userData){
			const data = JSON.parse(userData)

			if(data && data.token){
				login(data.token, data.userId)
			}
			setReady(true)
		}
		

	}, [login])


	return {logout, login, token, userId, ready}
}